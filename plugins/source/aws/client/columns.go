package client

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func DefaultAccountIDColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "account_id",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSAccount,
		PrimaryKeyComponent: pk,
	}
}

func RequestAccountIDColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "request_account_id",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSAccount,
		PrimaryKeyComponent: pk,
	}
}

func DefaultRegionColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "region",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSRegion,
		PrimaryKeyComponent: pk,
	}
}

func RequestRegionColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "request_region",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSRegion,
		PrimaryKeyComponent: pk,
	}
}

func LanguageCodeColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "language_code",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveLanguageCode,
		PrimaryKeyComponent: pk,
	}
}
