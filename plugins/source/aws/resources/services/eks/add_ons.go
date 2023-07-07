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

func addOns() *schema.Table {
	tableName := "aws_eks_cluster_addons"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Addon.html`,
		Resolver:            fetchAddOns,
		PreResourceResolver: getAddOn,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "eks"),
		Transform:           transformers.TransformWithStruct(&types.Addon{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AddonArn"),
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

func fetchAddOns(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Item.(*types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	paginator := eks.NewListAddonsPaginator(svc, &eks.ListAddonsInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Addons
	}
	return nil
}

func getAddOn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Eks
	name := resource.Item.(string)
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeAddon(
		ctx, &eks.DescribeAddonInput{
			ClusterName: cluster.Name,
			AddonName:   &name}, func(options *eks.Options) {
			options.Region = cl.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.Addon
	return nil
}
