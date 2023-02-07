package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFargateProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(*types.Cluster)
	c := meta.(*client.Client)
	svc := c.Services().Eks
	paginator := eks.NewListFargateProfilesPaginator(svc, &eks.ListFargateProfilesInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.FargateProfileNames
	}
	return nil
}

func getFargateProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Eks
	name := resource.Item.(string)
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeFargateProfile(
		ctx, &eks.DescribeFargateProfileInput{
			ClusterName:        cluster.Name,
			FargateProfileName: &name})
	if err != nil {
		return err
	}
	resource.Item = output.FargateProfile
	return nil
}
