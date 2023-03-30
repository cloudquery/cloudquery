package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func clusterParameterGroups() *schema.Table {
	tableName := "aws_redshift_cluster_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterParameterGroupStatus.html`,
		Resolver:    fetchClusterParameterGroups,
		Transform:   client.TransformWithStruct(&types.ClusterParameterGroupStatus{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
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
			clusterParameters(),
		},
	}
}

func fetchClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)
	res <- cluster.ClusterParameterGroups
	return nil
}
