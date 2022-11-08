package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func KafkaResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "endpoints",
			Struct:              &types.Cluster{},
			Description:         "https://docs.aws.amazon.com/MSK/2.0/APIReference/v2-clusters-clusterarn.html#v2-clusters-clusterarn-properties",
			SkipFields:          []string{"ClusterArn"},
			PreResourceResolver: "getCluster",
			ExtraColumns: append(defaultAccountColumns, []codegen.ColumnDefinition{{
				Name:     "arn",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("ClusterArn")`,
			},
			}...),
			ShouldGenerateResolverAndMockTest: true,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "kafka"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("kafka")`

		// Parameters for autogenerating the resolver and mock-test.
		// Only used when `ShouldGenerateResolverAndMockTest = true`
		r.ResolverAndMockTestTemplate = "list_and_describe_resources_1"
		r.CloudqueryServiceName = "Kafka"
	}

	return resources
}
