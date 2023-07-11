package eks

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func identityProviderConfigs() *schema.Table {
	tableName := "aws_eks_cluster_oidc_identity_provider_configs"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_OidcIdentityProviderConfig.html`,
		Resolver:            fetchIdentityProviderConfigs,
		PreResourceResolver: getIdentityProviderConfigs,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "eks"),
		Transform:           transformers.TransformWithStruct(&types.OidcIdentityProviderConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("IdentityProviderConfigArn"),
				PrimaryKey: true,
			},
			{
				Name:       "cluster_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIdentityProviderConfigs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Item.(*types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	paginator := eks.NewListIdentityProviderConfigsPaginator(svc, &eks.ListIdentityProviderConfigsInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.IdentityProviderConfigs
	}
	return nil
}

func getIdentityProviderConfigs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	ipc := resource.Item.(types.IdentityProviderConfig)
	if aws.ToString(ipc.Type) != "oidc" {
		return nil
	}
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeIdentityProviderConfig(
		ctx, &eks.DescribeIdentityProviderConfigInput{
			ClusterName:            cluster.Name,
			IdentityProviderConfig: &ipc}, func(options *eks.Options) {
			options.Region = cl.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.IdentityProviderConfig.Oidc
	return nil
}
