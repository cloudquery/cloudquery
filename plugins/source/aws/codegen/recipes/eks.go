package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EKSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "clusters",
			Struct:              &types.Cluster{},
			Description:         "https://docs.aws.amazon.com/eks/latest/APIReference/API_Cluster.html",
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getEksCluster",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "eks"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("eks")`
	}
	return resources
}
