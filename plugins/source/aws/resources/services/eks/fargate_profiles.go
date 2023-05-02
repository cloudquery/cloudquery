package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func fargateProfiles() *schema.Table {
	tableName := "aws_eks_fargate_profiles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_FargateProfile.html`,
		Resolver:            fetchFargateProfiles,
		PreResourceResolver: getFargateProfile,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "eks"),
		Transform:           transformers.TransformWithStruct(&types.FargateProfile{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FargateProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchFargateProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(*types.Cluster)
	c := meta.(*client.Client)
	svc := c.Services().Eks
	paginator := eks.NewListFargateProfilesPaginator(svc, &eks.ListFargateProfilesInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = c.Region
		})
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
			FargateProfileName: &name},
		func(options *eks.Options) {
			options.Region = c.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.FargateProfile
	return nil
}
