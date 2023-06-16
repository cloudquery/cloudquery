package client

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"google.golang.org/api/iterator"
)

func Identifier(name string) string {
	return "`" + name + "`"
}

type Stringer interface {
	String() string
}

func (*Client) listTables(ctx context.Context, client *firestore.Client) (schema.Tables, error) {
	var schemaTables schema.Tables
	collections := client.Collections(ctx)
	for {
		collection, err := collections.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}

		schemaTables = append(schemaTables, &schema.Table{
			Name: collection.ID,
			Columns: schema.ColumnList{
				{
					Name:       "__id",
					Type:       arrow.BinaryTypes.String,
					PrimaryKey: true,
					Unique:     true,
					NotNull:    true,
				},
				{Name: "__created_at", Type: arrow.FixedWidthTypes.Timestamp_us},
				{Name: "__updated_at", Type: arrow.FixedWidthTypes.Timestamp_us},
				{Name: "data", Type: types.ExtensionTypes.JSON},
			},
		})
	}
	return schemaTables, nil
}
