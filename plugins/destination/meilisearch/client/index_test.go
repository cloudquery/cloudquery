package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestMigrateAddColumnUpdatesFilterableAndSortableAttributes(t *testing.T) {
	ctx := context.Background()
	specBytes, err := json.Marshal(getTestSpec())
	require.NoError(t, err)
	pluginClient, err := New(ctx, zerolog.Nop(), specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pluginClient.(*Client)
	t.Cleanup(func() { require.NoError(t, c.Close(ctx)) })

	table := &schema.Table{
		Name:    fmt.Sprintf("test_migrate_add_column_%d", time.Now().UnixNano()),
		Columns: schema.ColumnList{{Name: "id", Type: arrow.BinaryTypes.String}},
	}
	t.Cleanup(func() { require.NoError(t, c.deleteIndex(ctx, tableIndexSchema(table))) })
	require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))

	table.Columns = append(table.Columns, schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String})
	require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))

	index := c.Meilisearch.Index(table.Name)
	filterable, err := index.GetFilterableAttributes()
	require.NoError(t, err)
	require.ElementsMatch(t, []any{"id", "policy_id"}, *filterable)

	sortable, err := index.GetSortableAttributes()
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"id", "policy_id"}, *sortable)
}
