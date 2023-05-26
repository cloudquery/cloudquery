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
		columns := make(schema.ColumnList, 0, 4)
		columns = append(columns, schema.Column{
			Name:       "__id",
			Type:       arrow.BinaryTypes.String,
			PrimaryKey: true,
			Unique:     true,
			NotNull:    true,
		})
		columns = append(columns, schema.Column{
			Name: "__created_at",
			Type: arrow.FixedWidthTypes.Timestamp_us,
		})
		columns = append(columns, schema.Column{
			Name: "__updated_at",
			Type: arrow.FixedWidthTypes.Timestamp_us,
		})
		columns = append(columns, schema.Column{
			Name: "data",
			Type: types.ExtensionTypes.JSON,
		})

		schemaTables = append(schemaTables, &schema.Table{
			Name:    collection.ID,
			Columns: columns,
		})
	}
	return schemaTables, nil
}
