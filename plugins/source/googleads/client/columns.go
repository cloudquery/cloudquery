package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var CustomerID = schema.Column{
	Name:            "customer_id",
	Type:            schema.TypeInt,
	Description:     "Customer ID",
	Resolver:        ResolveCustomerID,
	CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true, NotNull: true},
}

func IDColumn(path string) schema.Column {
	return schema.Column{
		Name:            "id",
		Type:            schema.TypeInt,
		Resolver:        schema.PathResolver(path),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true, NotNull: true},
	}
}
