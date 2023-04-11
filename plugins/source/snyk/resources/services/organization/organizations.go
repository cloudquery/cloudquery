package organization

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "snyk_organizations",
		Description: `https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Organization`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.SingleOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Organization{}, transformers.WithPrimaryKeys("ID")),
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
