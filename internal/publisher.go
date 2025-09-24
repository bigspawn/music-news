package internal

import (
	"context"
	"errors"
	"time"

	"github.com/go-pkgz/lgr"
	tb "gopkg.in/telebot.v3"
)

type PublisherParams struct {
	Lgr    lgr.L
	NewsCh <-chan []News
	BotAPI *RetryableBotApi
	Store  *Store
}

func (p *PublisherParams) Validate() error {
	if p.Lgr == nil {
		return errors.New("lgr is required")
	}
	if p.NewsCh == nil {
		return errors.New("news channel is required")
	}
	if p.BotAPI == nil {
		return errors.New("bot api is required")
	}
	if p.Store == nil {
		return errors.New("store is required")
	}
	return nil
}

type Publisher struct {
	PublisherParams
}

func NewPublisher(params PublisherParams) (*Publisher, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &Publisher{PublisherParams: params}, nil
}

func (p *Publisher) Start(ctx context.Context) error {
	p.Lgr.Logf("[INFO] Publisher started and waiting for news items")
	itemCount := 0
	startTime := time.Now()

	// Monitor channel every 5 minutes
	monitorTicker := time.NewTicker(5 * time.Minute)
	defer monitorTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			p.Lgr.Logf("[INFO] Publisher stopped after %v, total items processed: %d",
				time.Since(startTime), itemCount)
			return ctx.Err()
		case <-monitorTicker.C:
			p.Lgr.Logf("[INFO] Publisher monitor: running for %v, processed %d items, channel len=%d",
				time.Since(startTime), itemCount, len(p.NewsCh))
		case items := <-p.NewsCh:
			itemCount += len(items)
			p.Lgr.Logf("[INFO] Received batch of %d items, total processed: %d, channel len=%d",
				len(items), itemCount, len(p.NewsCh))

			err := p.publish(ctx, items)
			if err != nil {
				p.Lgr.Logf("[ERROR] publishing batch: %v", err)
			} else {
				p.Lgr.Logf("[INFO] Successfully processed batch of %d items", len(items))
			}
		}
	}
}

func (p *Publisher) Stop() {}

func (p *Publisher) publish(ctx context.Context, items []News) error {
	// Deduplication: filter out already sent items
	var filteredItems []News
	for _, item := range items {
		isPosted, err := p.Store.IsPosted(ctx, item.ID)
		if err != nil {
			p.Lgr.Logf("[ERROR] failed to check if item is posted: ID=%d, err=%v", item.ID, err)
			continue
		}
		if isPosted {
			p.Lgr.Logf("[DEBUG] skipping already posted item: ID=%d, Title=[%s]", item.ID, item.Title)
			continue
		}
		filteredItems = append(filteredItems, item)
	}

	if len(filteredItems) == 0 {
		p.Lgr.Logf("[INFO] All items in batch already posted, nothing to process")
		return nil
	}

	if len(filteredItems) < len(items) {
		p.Lgr.Logf("[INFO] Filtered out %d already posted items, processing %d remaining items",
			len(items)-len(filteredItems), len(filteredItems))
	}

	successCount := 0
	errorCount := 0

	for i, item := range filteredItems {
		p.Lgr.Logf("[DEBUG] Processing item %d/%d: ID=%d, Title=[%s]", i+1, len(filteredItems), item.ID, item.Title)

		if err := p.BotAPI.SendNews(ctx, item); err != nil {
			errorCount++
			p.Lgr.Logf("[ERROR] can't send news: item={ID=%d, Title=%s, ImageLink=%s}, err=%v",
				item.ID, item.Title, item.ImageLink, err)

			var bErr *tb.Error
			if errors.As(err, &bErr) && bErr.Code == 400 &&
				(bErr.Message == "wrong type of the web page content" ||
				 bErr.Message == "failed to get HTTP URL content") {
				p.Lgr.Logf("[INFO] blacklisting problematic image item: ID=%d, Title=[%s], Error=[%s]", item.ID, item.Title, bErr.Message)
				_ = p.Store.SetPostedAndNotified(ctx, item.ID)
			}

			continue
		}

		if err := p.Store.SetPostedByID(ctx, item.ID); err != nil {
			p.Lgr.Logf("[ERROR] can't set posted flag: item={ID=%d, Title=%s}, err=%v", item.ID, item.Title, err)
			continue
		}

		successCount++
		p.Lgr.Logf("[INFO] news successfully sent (%d/%d): [%s]", successCount, len(filteredItems), item.Title)

		// Add random delay only if not the last item
		if i < len(filteredItems)-1 {
			duration := time.Duration(RandBetween(10_000, 1)) * time.Millisecond
			p.Lgr.Logf("[DEBUG] sleeping %s before next send", duration)
			WaitUntil(ctx, duration)
		}
	}

	p.Lgr.Logf("[INFO] Batch processing completed: success=%d, errors=%d, total=%d",
		successCount, errorCount, len(filteredItems))

	return nil
}
