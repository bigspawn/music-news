package internal

import (
	"context"

	"github.com/mmcdole/gofeed"
)

type CoreRadioParser struct{}

func (p *CoreRadioParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
	//TODO implement me
	panic("implement me")
}
