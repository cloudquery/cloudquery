package client

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/azblob/v4/client/spec"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestTableNameOutsidePathIsRejected(t *testing.T) {
	ctx := context.Background()
	c := &Client{spec: &spec.Spec{Path: "cq/data", NoRotate: true, FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON}}}
	table := &schema.Table{Name: "../x", Columns: schema.ColumnList{{Name: "id", Type: arrow.BinaryTypes.String}}}

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	bldr.Field(0).(*array.StringBuilder).Append("value")
	msgs := make(chan *message.WriteInsert, 1)
	msgs <- &message.WriteInsert{Record: bldr.NewRecordBatch()}
	close(msgs)

	require.ErrorContains(t, c.WriteTable(ctx, msgs), "cannot be used in a path")
	require.ErrorContains(t, c.Read(ctx, table, nil), "cannot be used in a path")
}
