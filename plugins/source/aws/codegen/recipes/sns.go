package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SNSResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "subscriptions",
			Struct:     &sns.Subscription{},
			SkipFields: []string{"SubscriptionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("SubscriptionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},

		{
			SubService: "topics",
			Struct:     &sns.Topic{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveSnsTopicTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "sns"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("sns")`
	}
	return resources
}
