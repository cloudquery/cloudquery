package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:        "aws_organization_policies",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html`,
		Resolver:    fetchOrganizationsPolicies,
		Transform:   transformers.TransformWithStruct(&types.Policy{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			// This is needed as a PK because aws managed policies don't have an account_id in the ARN
			client.DefaultAccountIDColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicySummary.Arn"),
			},
		},
	}
}

func fetchOrganizationsPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	for _, policyType := range types.PolicyType("").Values() { {
		paginator := organizations.NewListPoliciesPaginator(c.Services().Organizations, &organizations.ListPoliciesInput{
			Filter: policyType,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return err
			}
			res <- page.Policies
		}
	}

	return nil
}
