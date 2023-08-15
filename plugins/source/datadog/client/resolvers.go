package client

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ResolveAccountName(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.multiplexedAccount.Name)
}

var AccountNameColumn = schema.Column{
	Name:       "account_name",
	Type:       arrow.BinaryTypes.String,
	Resolver:   ResolveAccountName,
	PrimaryKey: true,
	NotNull:    true,
}
