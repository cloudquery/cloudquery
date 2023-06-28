package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (*Client) DeleteStale(context.Context, <-chan *message.WriteDeleteStale) error {
	return fmt.Errorf("destination doesn't support overwrite-delete-stale mode. please use append mode")
}
