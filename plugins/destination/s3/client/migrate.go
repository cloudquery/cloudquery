package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/writers"
)

func (*Client) MigrateTable(_ context.Context, ch <-chan *message.WriteMigrateTable) error {
	// migrate is not needed in append mode
	writers.DummyHandler[*message.WriteMigrateTable](ch)
	return nil
}
