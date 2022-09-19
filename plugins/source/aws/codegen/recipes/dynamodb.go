package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DynamoDBResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "tables",
			Struct:     &types.TableDescription{},
			SkipFields: []string{"TableArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDynamodbTableTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TableArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"TableReplicaAutoScalings()",
				"TableContinuousBackups()",
			},
		},
		{
			SubService: "table_replica_auto_scalings",
			Struct:     &types.ReplicaAutoScalingDescription{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "table_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "table_continuous_backups",
			Struct:     &types.ContinuousBackupsDescription{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "table_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "dynamodb"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("dynamodb")`
	}
	return resources
}
