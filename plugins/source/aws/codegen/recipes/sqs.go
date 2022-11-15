package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SQSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "queues",
			Struct:              &models.Queue{},
			SkipFields:          []string{"Arn", "Policy", "RedriveAllowPolicy", "RedrivePolicy"},
			PreResourceResolver: "getQueue",
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
					{
						Name:     "policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("Policy")`,
					},
					{
						Name:     "redrive_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("RedrivePolicy")`,
					},
					{
						Name:     "redrive_allow_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("RedriveAllowPolicy")`,
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
