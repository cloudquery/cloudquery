// Code generated by codegen; DO NOT EDIT.

package redshift

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_redshift_cluster_parameter_groups",
		Resolver:  fetchRedshiftClusterParameterGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("redshift"),
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
				Name:        "cluster_arn",
				Type:        schema.TypeString,
				Resolver:    schema.ParentResourceFieldResolver("arn"),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "parameter_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterGroupName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cluster_parameter_status_list",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ClusterParameterStatusList"),
			},
			{
				Name:     "parameter_apply_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterApplyStatus"),
			},
		},

		Relations: []*schema.Table{
			ClusterParameters(),
		},
	}
}
