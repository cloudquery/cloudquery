package organization

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func pendingUserProvisions() *schema.Table {
	return &schema.Table{
		Name:        "snyk_organization_provisions",
		Description: `https://snyk.docs.apiary.io/#reference/organizations/provision-user/list-pending-user-provisions`,
		Resolver:    fetchPendingProvisions,
		Multiplex:   client.ByOrganization,
		Transform: transformers.TransformWithStruct(&snyk.PendingProvision{},
			transformers.WithPrimaryKeys("Email", "Created")),
		Columns: []schema.Column{
			{
				Name:       "organization_id",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
				Resolver:   schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchPendingProvisions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	opts := snyk.ListPendingUserProvisionsOptions{
		Page:    1,
		PerPage: 1000,
	}

	for {
		result, _, err := c.Orgs.ListPendingUserProvisions(ctx, c.OrganizationID, opts)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			break
		}
		res <- result
		opts.Page++
	}

	return nil
}
