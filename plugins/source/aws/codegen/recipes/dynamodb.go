package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DynamoDBResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "tables",
			Struct:              &types.TableDescription{},
			Description:         "https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html",
			SkipFields:          []string{"TableArn"},
			PreResourceResolver: "getTable",
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
			SubService:  "table_replica_auto_scalings",
			Struct:      &types.ReplicaAutoScalingDescription{},
			Description: "https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "table_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "table_continuous_backups",
			Struct:      &types.ContinuousBackupsDescription{},
			Description: "https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ContinuousBackupsDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "table_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
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
