// Code generated by codegen; DO NOT EDIT.

package neptune

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_neptune_cluster_parameter_groups",
		Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameters",
		Resolver:    fetchNeptuneClusterParameterGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("neptune"),
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
				Resolver: resolveNeptuneClusterParameterGroupTags,
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
