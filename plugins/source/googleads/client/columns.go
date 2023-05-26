package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var CustomerID = schema.Column{
	Name:        "customer_id",
	Type:        arrow.PrimitiveTypes.Int64,
	Description: "Customer ID",
	PrimaryKey:  true,
	NotNull:     true,
	Resolver:    ResolveCustomerID,
}

func IDColumn(path string) schema.Column {
	return schema.Column{
		Name:       "id",
		Type:       arrow.PrimitiveTypes.Int64,
		PrimaryKey: true,
		NotNull:    true,
		Resolver:   schema.PathResolver(path),
	}
}
