package organizations

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Policies() *schema.Table {
	tableName := "aws_organizations_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html`,
		Resolver:    fetchOrganizationsPolicies,
		Transform:   transformers.TransformWithStruct(&types.PolicySummary{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			// This is needed as a PK because aws managed policies don't have an account_id in the ARN
			client.DefaultAccountIDColumn(true),
			{
				Name:     "content",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolvePolicyContent,
			},
		},
	}
}

func fetchOrganizationsPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Organizations
	for _, policyType := range types.PolicyType("").Values() {
		paginator := organizations.NewListPoliciesPaginator(svc, &organizations.ListPoliciesInput{
			Filter: policyType,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- page.Policies
		}
	}
	return nil
}

func resolvePolicyContent(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicySummary)
	cl := meta.(*client.Client)
	svc := cl.Services().Organizations
	resp, err := svc.DescribePolicy(ctx, &organizations.DescribePolicyInput{
		PolicyId: r.Id,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, resp.Policy.Content)
}
