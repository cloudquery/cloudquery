package eks

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func nodeGroups() *schema.Table {
	tableName := "aws_eks_cluster_node_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html`,
		Resolver:            fetchNodeGroups,
		PreResourceResolver: getNodeGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "eks"),
		Transform:           transformers.TransformWithStruct(&types.Nodegroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("NodegroupArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchNodeGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Item.(*types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	paginator := eks.NewListNodegroupsPaginator(svc, &eks.ListNodegroupsInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Nodegroups
	}
	return nil
}

func getNodeGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	name := resource.Item.(string)
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeNodegroup(
		ctx, &eks.DescribeNodegroupInput{
			ClusterName:   cluster.Name,
			NodegroupName: &name}, func(options *eks.Options) {
			options.Region = cl.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.Nodegroup
	return nil
}
