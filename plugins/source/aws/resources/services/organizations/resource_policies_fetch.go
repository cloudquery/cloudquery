package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchOrganizationsResourcePolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	o, err := c.Services().Organizations.DescribeResourcePolicy(ctx, &organizations.DescribeResourcePolicyInput{})
	if err != nil {
		return err
	}

	res <- o.ResourcePolicy
	return nil
}
