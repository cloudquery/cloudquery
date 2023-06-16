package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (*Client) Read(context.Context, *schema.Table, string, chan<- arrow.Record) error {
	return fmt.Errorf("read is not implemented")
}
