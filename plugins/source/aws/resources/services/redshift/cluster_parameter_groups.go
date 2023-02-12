package redshift

import (
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_cluster_parameter_groups",
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterParameterGroupStatus.html`,
		Resolver:    fetchRedshiftClusterParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.ClusterParameterGroupStatus{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "cluster_arn",
				Type:        schema.TypeString,
				Resolver:    schema.ParentColumnResolver("arn"),
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
		},

		Relations: []*schema.Table{
			ClusterParameters(),
		},
	}
}
