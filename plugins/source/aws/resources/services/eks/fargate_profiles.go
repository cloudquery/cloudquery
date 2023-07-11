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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FargateProfileArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchFargateProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(*types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	paginator := eks.NewListFargateProfilesPaginator(svc, &eks.ListFargateProfilesInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.FargateProfileNames
	}
	return nil
}

func getFargateProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	name := resource.Item.(string)
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeFargateProfile(
		ctx, &eks.DescribeFargateProfileInput{
			ClusterName:        cluster.Name,
			FargateProfileName: &name},
		func(options *eks.Options) {
			options.Region = cl.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.FargateProfile
	return nil
}
