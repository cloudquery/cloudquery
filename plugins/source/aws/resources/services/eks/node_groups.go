package eks

import (
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func NodeGroups() *schema.Table {
	return &schema.Table{
		Name:                "aws_eks_cluster_node_groups",
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html`,
		Resolver:            fetchNodeGroups,
		PreResourceResolver: getNodeGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer("eks"),
		Transform:           transformers.TransformWithStruct(&types.Nodegroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodegroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
