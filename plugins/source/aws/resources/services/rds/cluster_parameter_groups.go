// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_parameter_groups",
		Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterParameterGroup.html",
		Resolver:    fetchRdsClusterParameterGroups,
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
			{
				Name:     "db_cluster_parameter_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroupName"),
			},
			{
				Name:     "db_parameter_group_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupFamily"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
		},

		Relations: []*schema.Table{
			ClusterParameterGroupParameters(),
		},
	}
}
