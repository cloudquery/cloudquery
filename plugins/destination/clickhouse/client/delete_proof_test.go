package client

import (
	"context"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProof(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()

	var messages []message.WriteMessage

	p := initPlugin(t, &spec.Spec{})

	table := createTestTable()
	messages = append(messages, &message.WriteMigrateTable{Table: table})
	messages = append(messages, createInsertMessages([]string{"foo"}, table)[0])
	messages = append(messages, createDeleteMessages(false, table, []string{"foo"})[0])
	messages = append(messages, createInsertMessages([]string{"bar"}, table)[0])
	r.NoError(p.WriteAll(ctx, messages))

	count, err := countAllRowsFromPlugin(ctx, p, table)
	r.NoError(err)
	r.EqualValues(int64(1), count, "unexpected amount of items after delete with match")
}

func countAllRowsFromPlugin(ctx context.Context, p *plugin.Plugin, table *schema.Table) (int64, error) {
	var err error
	ch := make(chan arrow.Record)
	go func() {
		defer close(ch)
		err = p.Read(ctx, table, ch)
	}()
	count := int64(0)
	for record := range ch {
		count += record.NumRows()
	}
	return count, err
}
