package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SQSResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "queues",
			Struct:     &sqs.Queue{},
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
						Resolver: `resolveSqsQueueTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "sqs"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("sqs")`
	}
	return resources
}
