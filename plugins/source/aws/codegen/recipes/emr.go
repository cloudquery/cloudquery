package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EMRResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "block_public_access_configs",
			Struct:     &emr.GetBlockPublicAccessConfigurationOutput{},
			SkipFields: []string{"_"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:          "clusters",
			Struct:              &types.Cluster{},
			Description:         "https://docs.aws.amazon.com/emr/latest/APIReference/API_Cluster.html",
			SkipFields:          []string{"ClusterArn"},
			PreResourceResolver: "getCluster",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ClusterArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "emr"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticmapreduce")`
	}
	return resources
}
