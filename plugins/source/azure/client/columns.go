package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
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
