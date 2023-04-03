package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func tableContinuousBackups() *schema.Table {
	tableName := "aws_dynamodb_table_continuous_backups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ContinuousBackupsDescription.html`,
		Resolver:    fetchDynamodbTableContinuousBackups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:   transformers.TransformWithStruct(&types.ContinuousBackupsDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "table_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
