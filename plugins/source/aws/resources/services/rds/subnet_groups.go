package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_subnet_groups",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSubnetGroup.html`,
		Resolver:    fetchRdsSubnetGroups,
		Transform:   transformers.TransformWithStruct(&types.DBSubnetGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
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
				Resolver: schema.PathResolver("DBSubnetGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
