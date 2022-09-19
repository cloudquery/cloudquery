package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DaxResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "clusters",
			Struct:     &types.Cluster{},
			SkipFields: []string{"ClusterArn"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("dax")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ClusterArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeString,
						Resolver: `resolveClusterTags`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "dax"
	}
	return resources
}
