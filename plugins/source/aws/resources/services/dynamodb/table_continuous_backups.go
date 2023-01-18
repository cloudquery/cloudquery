package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TableContinuousBackups() *schema.Table {
	return &schema.Table{
		Name:        "aws_dynamodb_table_continuous_backups",
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ContinuousBackupsDescription.html`,
		Resolver:    fetchDynamodbTableContinuousBackups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dynamodb"),
		Transform:   transformers.TransformWithStruct(&types.ContinuousBackupsDescription{}),
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
				Name:     "table_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
