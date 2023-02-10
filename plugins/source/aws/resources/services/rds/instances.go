package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_instances",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html`,
		Resolver:    fetchRdsInstances,
		Transform:   transformers.TransformWithStruct(&types.DBInstance{}, transformers.WithSkipFields("TagList")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.AccountPKColumn(false),
			client.RegionPKColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "processor_features",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstanceProcessorFeatures,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstanceTags,
			},
		},
	}
}
