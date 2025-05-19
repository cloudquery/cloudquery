package client

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/stretchr/testify/require"
)

func TestDeleteAfterInsert(t *testing.T) {
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

	count, err := countAllRows(ctx, p, table)
	r.NoError(err)
	r.EqualValues(int64(1), count, "unexpected amount of items after delete with match")
}
