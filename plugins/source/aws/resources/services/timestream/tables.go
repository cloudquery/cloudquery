package timestream

import (
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:        "aws_timestream_tables",
		Description: `https://docs.aws.amazon.com/timestream/latest/developerguide/API_Table.html`,
		Resolver:    fetchTimestreamTables,
		Transform:   transformers.TransformWithStruct(&types.Table{}),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
