package client

import (
	"context"

	"github.com/apache/arrow/go/v12/arrow"
)

func (*Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	panic("not implemented")
}
