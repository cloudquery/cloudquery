package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudHSMV2() []*Resource {
	resources := []*Resource{
		{
			SubService: "clusters",
			Struct:     &types.Cluster{},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("cloudhsmv2")`,
			SkipFields: []string{"TagList"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveClusterArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTagField("TagList")`,
					},
				}...),
		},
		{
			SubService: "backups",
			Struct:     &types.Backup{},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("cloudhsmv2")`,
			SkipFields: []string{"TagList"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveBackupArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					}, {
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTagField("TagList")`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "cloudhsmv2"
	}
	return resources
}
