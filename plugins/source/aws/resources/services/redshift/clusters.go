package redshift

import (
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_clusters",
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchRedshiftClusters,
		Transform:   transformers.TransformWithStruct(&types.Cluster{}, transformers.WithSkipFields("ClusterParameterGroups")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveClusterArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "logging_status",
				Type:        schema.TypeJSON,
				Resolver:    resolveRedshiftClusterLoggingStatus,
				Description: `Describes the status of logging for a cluster.`,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			Snapshots(),
			ClusterParameterGroups(),
		},
	}
}
