package xray

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ResourcePolicies() *schema.Table {
	tableName := "aws_xray_resource_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_ResourcePolicy.html`,
		Resolver:    fetchXrayResourcePolicies,
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "policy_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PolicyName"),
				PrimaryKey: true,
			},
			{
				Name:       "policy_revision_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PolicyRevisionId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchXrayResourcePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := xray.NewListResourcePoliciesPaginator(cl.Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(o *xray.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.ResourcePolicies
	}
	return nil
}
