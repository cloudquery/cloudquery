package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	IDColumn = schema.Column{
		Name:            "id",
		Type:            schema.TypeString,
		Resolver:        schema.PathResolver("ID"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	SubscriptionID = schema.Column{
		Name:     "subscription_id",
		Type:     schema.TypeString,
		Resolver: ResolveAzureSubscription,
	}
	SubscriptionIDPK = schema.Column{
		Name:            "subscription_id",
		Type:            schema.TypeString,
		Resolver:        ResolveAzureSubscription,
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
)
