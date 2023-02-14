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
			client.DefaultAccountIDColumn(false),
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

	for _, policy_type := range []types.PolicyType{types.PolicyTypeServiceControlPolicy, types.PolicyTypeTagPolicy, types.PolicyTypeBackupPolicy, types.PolicyTypeAiservicesOptOutPolicy} {
		paginator := organizations.NewListPoliciesPaginator(c.Services().Organizations, &organizations.ListPoliciesInput{
			Filter: policy_type,
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
