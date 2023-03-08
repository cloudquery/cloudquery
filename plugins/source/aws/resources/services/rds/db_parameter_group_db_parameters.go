package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DbParameterGroupDbParameters() *schema.Table {
	tableName := "aws_rds_db_parameter_group_db_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html`,
		Resolver:    fetchRdsDbParameterGroupDbParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "db_parameter_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
