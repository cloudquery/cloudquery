package eks

import (
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:                "aws_eks_clusters",
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Cluster.html`,
		Resolver:            fetchEksClusters,
		PreResourceResolver: getEksCluster,
		Multiplex:           client.ServiceAccountRegionMultiplexer("eks"),
		Transform:           transformers.TransformWithStruct(&types.Cluster{}),
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
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
