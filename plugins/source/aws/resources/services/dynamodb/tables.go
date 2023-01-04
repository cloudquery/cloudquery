package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "aws_dynamodb_tables",
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html`,
		Resolver:            fetchDynamodbTables,
		PreResourceResolver: getTable,
		Multiplex:           client.ServiceAccountRegionMultiplexer("dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.TableDescription{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDynamodbTableTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TableArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "archival_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ArchivalSummary"),
			},
		},
		Relations: []*schema.Table{
			TableReplicaAutoScalings(),
			TableContinuousBackups(),
		},
	}
}
