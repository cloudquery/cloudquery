package cloudhsmv2

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudhsmv2_clusters",
		Description: `https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchCloudhsmv2Clusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudhsmv2"),
		Transform:   transformers.TransformWithStruct(&types.Cluster{}, transformers.WithSkipFields("TagList")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveClusterArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
		},
	}
}
