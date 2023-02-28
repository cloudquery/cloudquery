package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourcePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_organization_resource_policies",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DescribeResourcePolicy.html`,
		Resolver:    fetchOrganizationsResourcePolicies,
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("organizations"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchOrganizationsResourcePolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	o, err := c.Services().Organizations.DescribeResourcePolicy(ctx, &organizations.DescribeResourcePolicyInput{})
	if err != nil {
		return err
	}

	res <- o.ResourcePolicy
	return nil
}
