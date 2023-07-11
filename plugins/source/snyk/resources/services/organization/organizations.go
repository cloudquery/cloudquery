package organization

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "snyk_organizations",
		Description: `https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.SingleOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Organization{}, transformers.WithPrimaryKeys("ID")),
		Relations: []*schema.Table{
			organizationMembers(),
			pendingUserProvisions(),
		},
	}
}

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, _, err := c.Orgs.List(ctx)
	if err != nil {
		return err
	}

	// limit organizations
	for _, org := range result {
		if c.WantOrganization(org.ID) {
			res <- org
		}
	}

	return nil
}
