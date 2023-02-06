package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchNodeGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Item.(*types.Cluster)
	c := meta.(*client.Client)
	svc := c.Services().Eks
	paginator := eks.NewListNodegroupsPaginator(svc, &eks.ListNodegroupsInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Nodegroups
	}
	return nil
}

func getNodeGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Eks
	name := resource.Item.(string)
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeNodegroup(
		ctx, &eks.DescribeNodegroupInput{
			ClusterName:   cluster.Name,
			NodegroupName: &name})
	if err != nil {
		return err
	}
	resource.Item = output.Nodegroup
	return nil
}
