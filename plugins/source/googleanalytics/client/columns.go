package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var PropertyIDColumn = schema.Column{
	Name:        "property_id",
	Type:        schema.TypeString,
	Description: "Property ID",
	Resolver: func(_ context.Context, m schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, m.(*Client).PropertyID)
	},
	CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true, NotNull: true},
}

var DateColumn = schema.Column{
	Name:        "date",
	Type:        schema.TypeTimestamp,
	Description: "Date",
	Resolver:    schema.PathResolver("Date"),
	CreationOptions: schema.ColumnCreationOptions{
		PrimaryKey:     true,
		IncrementalKey: true,
		NotNull:        true,
	},
}
