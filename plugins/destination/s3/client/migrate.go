package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/writers/streamingbatchwriter"
)

func (*Client) MigrateTable(_ context.Context, ch <-chan *message.WriteMigrateTable) error {
	// migrate is not needed in append mode
	streamingbatchwriter.DummyHandler[*message.WriteMigrateTable](ch)
	return nil
}
