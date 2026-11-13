package client

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestMigrateTwiceTableNameWithQuoteOrBackslash(t *testing.T) {
	ctx := context.Background()
	for _, name := range []string{"qa_plain_ttl", `qa"dq_ttl`, `qa\back_ttl`, "qa`tick_ttl"} {
		t.Run(name, func(t *testing.T) {
			p := initPlugin(t)
			table := &schema.Table{Name: name, Columns: schema.ColumnList{
				{Name: "_cq_id", Type: arrow.BinaryTypes.String, NotNull: true, PrimaryKey: true},
			}}
			messages := []message.WriteMessage{&message.WriteMigrateTable{Table: table}}
			require.NoError(t, p.WriteAll(ctx, messages))
			require.NoError(t, p.WriteAll(ctx, messages))
		})
	}
}
