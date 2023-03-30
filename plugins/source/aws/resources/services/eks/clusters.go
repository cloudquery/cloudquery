package eks

import (
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	tableName := "aws_eks_clusters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Cluster.html`,
		Resolver:            fetchEksClusters,
		PreResourceResolver: getEksCluster,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "eks"),
		Transform:           client.TransformWithStruct(&types.Cluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			NodeGroups(),
			FargateProfiles(),
		},
	}
}
