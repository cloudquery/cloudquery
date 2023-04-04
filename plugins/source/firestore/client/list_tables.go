package client

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/iterator"
)

func Identifier(name string) string {
	return "`" + name + "`"
}

type Stringer interface {
	String() string
}

func (c *Client) listTables(ctx context.Context, client *firestore.Client) (schema.Tables, error) {
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
			Name: "_id",
			Type: schema.TypeString,
			CreationOptions: schema.ColumnCreationOptions{
				PrimaryKey: true,
				NotNull:    true,
				Unique:     true,
			},
		})
		columns = append(columns, schema.Column{
			Name: "_created_at",
			Type: schema.TypeTimestamp,
		})
		columns = append(columns, schema.Column{
			Name: "_updated_at",
			Type: schema.TypeTimestamp,
		})
		columns = append(columns, schema.Column{
			Name: "data",
			Type: schema.TypeJSON,
		})

		schemaTables = append(schemaTables, &schema.Table{
			Name:    collection.ID,
			Columns: columns,
		})
	}
	return schemaTables, nil
}
