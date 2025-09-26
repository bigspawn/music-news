package internal

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-pkgz/lgr"
)

const (
	driver      = "sqlite3"
	insertQuery = `			insert into main.news(title, playlist, date_time, imageurl, downloadurl, pageurl, posted, created_at)
							values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	selectNotified = `		select id, title, playlist, imageurl, date_time, downloadurl
							from main.news
							where not notified
							  and title not like '%Single%'
							  and title not like '%single%'
							  and created_at > strftime('%Y-%m-%d %H:%M:%S+00:00', 'now', 'utc', '-1 months')
							order by date_time`

	updateNotified = `		update main.news
							set notified = true
							where id = $1`

	selectUnpublished = `	select id, title, playlist, imageurl, date_time, downloadurl, pageurl
							from main.news
							where posted = false
							  and created_at > strftime('%Y-%m-%d %H:%M:%S+00:00', 'now', 'utc', '-1 months')
							order by created_at`

	updatePosted = `		update main.news
							set posted = true
							where title = $1`

	selectExists = `		select *
							from main.news
							where title like '%' || $1 || '%'`

	selectAll = `			select id, title, date_time, downloadurl, imageurl, playlist, posted, notified, created_at
							from main.news`

	updatePostedByID = `	update main.news
							set posted = true
							where id = $1`

	updatePostedAndNotifiedByID = `	update main.news
									set posted = true, notified = true
									where id = $1`
)

// News is an article structure.
type News struct {
	ID           int
	Title        string
	DateTime     time.Time
	DownloadLink []string
	ImageLink    string
	PageLink     string
	Text         string
	Posted       bool
	Notified     bool
	CreatedAt    time.Time
	TitleHash    string
}

type StoreParams struct {
	Lgr lgr.L
	DB  *sql.DB
}

func (p *StoreParams) Validate() error {
	if p.DB == nil {
		return fmt.Errorf("db is required")
	}
	if p.Lgr == nil {
		return fmt.Errorf("lgr is required")
	}
	return nil
}

type Store struct {
	StoreParams
}

func NewStore(params StoreParams) (*Store, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &Store{StoreParams: params}, nil
}

func (s *Store) Exist(ctx context.Context, title string) (bool, error) {
	row, err := s.DB.QueryContext(ctx, selectExists, title)
	if err != nil {
		return false, err
	}
	defer func() { _ = row.Close() }()

	if row != nil && row.Next() {
		return true, nil
	}
	return false, nil
}

func (s *Store) Insert(ctx context.Context, n News) (int, error) {
	link := ""
	if len(n.DownloadLink) > 0 {
		link = n.DownloadLink[0]
	}
	var id int
	row := s.DB.QueryRowContext(
		ctx,
		insertQuery,
		n.Title,
		n.Text,
		n.DateTime,
		n.ImageLink,
		link,
		n.PageLink,
		false,
		time.Now().UTC(),
	)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Store) GetWithNotifyFlag(ctx context.Context) ([]News, error) {
	rows, err := s.DB.QueryContext(ctx, selectNotified)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var (
		id           = new(int)
		title        = new(string)
		text         = new(string)
		imageLink    = new(string)
		downloadLink = new(string)
		dateTime     = new(time.Time)
		notifies     = make([]News, 0)
	)
	for rows.Next() {
		if err := rows.Scan(id, title, text, imageLink, dateTime, downloadLink); err != nil {
			return nil, err
		}
		notifies = append(notifies, News{
			ID:           *id,
			Title:        *title,
			Text:         *text,
			ImageLink:    *imageLink,
			DownloadLink: strings.Split(*downloadLink, ","),
			DateTime:     *dateTime,
		})
	}
	return notifies, nil
}

func (s *Store) UpdateNotifyFlag(ctx context.Context, item News) error {
	return s.exec(ctx, updateNotified, item.ID)
}

func (s *Store) GetUnpublished(ctx context.Context) ([]News, error) {
	rows, err := s.DB.QueryContext(ctx, selectUnpublished)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var (
		ID           = new(int)
		Title        = new(string)
		Text         = new(string)
		ImageLink    = new(string)
		PageLink     = new(string)
		DownloadLink = new(string)
		DateTime     = new(time.Time)
		unpublished  = make([]News, 0, 32)
	)
	for rows.Next() {
		if err := rows.Scan(ID, Title, Text, ImageLink, DateTime, DownloadLink, PageLink); err != nil {
			return nil, err
		}
		n := News{
			ID:           *ID,
			Title:        *Title,
			Text:         *Text,
			ImageLink:    *ImageLink,
			DownloadLink: strings.Split(*DownloadLink, ","),
			DateTime:     *DateTime,
			PageLink:     *PageLink,
		}
		unpublished = append(unpublished, n)
	}
	return unpublished, nil
}

func (s *Store) SetPosted(ctx context.Context, title string) error {
	return s.exec(ctx, updatePosted, title)
}

func (s *Store) SetPostedByID(ctx context.Context, id int) error {
	return s.exec(ctx, updatePostedByID, id)
}

func (s *Store) GetAll(ctx context.Context) ([]News, error) {
	rows, err := s.DB.QueryContext(ctx, selectAll)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	items := make([]News, 0)
	for rows.Next() {
		var n News
		var links string
		err := rows.Scan(
			&n.ID,
			&n.Title,
			&n.DateTime,
			&links,
			&n.ImageLink,
			&n.Text,
			&n.Posted,
			&n.Notified,
			&n.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		n.DownloadLink = strings.Split(links, ",")
		items = append(items, n)
	}

	return items, nil
}

func (s *Store) SetPostedAndNotified(ctx context.Context, id int) error {
	return s.exec(ctx, updatePostedAndNotifiedByID, id)
}

func (s *Store) Stop() {
	if err := s.DB.Close(); err != nil {
		s.Lgr.Logf("[ERROR] failed to close db connection: %v", err)
	}
}

func (s *Store) exec(ctx context.Context, sql string, args ...interface{}) error {
	result, err := s.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}
