package client

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var PropertyIDColumn = schema.Column{
	Name:        "property_id",
	Type:        arrow.BinaryTypes.String,
	Description: "Property ID",
	PrimaryKey:  true,
	NotNull:     true,
	Resolver: func(_ context.Context, m schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, m.(*Client).PropertyID)
	},
}

var DateColumn = schema.Column{
	Name:           "date",
	Type:           arrow.FixedWidthTypes.Timestamp_us,
	Description:    "Date",
	PrimaryKey:     true,
	IncrementalKey: true,
	NotNull:        true,
	Resolver:       schema.PathResolver("Date"),
}
