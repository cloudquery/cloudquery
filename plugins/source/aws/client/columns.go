package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func DefaultAccountIDColumn(pk bool) schema.Column {
	return schema.Column{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: ResolveAWSAccount,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: pk,
		},
	}
}

func DefaultRegionColumn(pk bool) schema.Column {
	return schema.Column{
		Name:     "region",
		Type:     schema.TypeString,
		Resolver: ResolveAWSRegion,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: pk,
		},
	}
}

func LanguageCodeColumn(pk bool) schema.Column {
	return schema.Column{
		Name:     "language_code",
		Type:     schema.TypeString,
		Resolver: ResolveLanguageCode,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: pk,
		},
	}
}
