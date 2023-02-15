package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
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
