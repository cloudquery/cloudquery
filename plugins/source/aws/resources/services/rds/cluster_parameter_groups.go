package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_parameter_groups",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterParameterGroup.html`,
		Resolver:    fetchRdsClusterParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.DBClusterParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsClusterParameterGroupTags,
			},
		},

		Relations: []*schema.Table{
			ClusterParameterGroupParameters(),
		},
	}
}
