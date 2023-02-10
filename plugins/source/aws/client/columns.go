package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountPKColumn(pk bool) schema.Column {
	return schema.Column{
		Name:        "account_id",
		Type:        schema.TypeString,
		Resolver:    ResolveAWSAccount,
		RetainOrder: true,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: pk,
		},
	}
}

func RegionPKColumn(pk bool) schema.Column {
	return schema.Column{
		Name:        "region",
		Type:        schema.TypeString,
		Resolver:    ResolveAWSRegion,
		RetainOrder: true,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: pk,
		},
	}
}
