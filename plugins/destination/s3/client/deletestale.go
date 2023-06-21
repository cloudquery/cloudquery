package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (*Client) DeleteStale(context.Context, []*message.DeleteStale) error {
	return fmt.Errorf("destination plugin doesn't support overwrite-delete-stale mode. please use append mode")
}
