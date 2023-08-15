package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	SubscriptionID = schema.Column{
		Name:     "subscription_id",
		Type:     arrow.BinaryTypes.String,
		Resolver: ResolveAzureSubscription,
	}
	SubscriptionIDPK = schema.Column{
		Name:       "subscription_id",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveAzureSubscription,
		PrimaryKey: true,
	}
)
