package organization

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func organizationMembers() *schema.Table {
	return &schema.Table{
		Name:        "snyk_organization_members",
		Description: `https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members`,
		Resolver:    fetchOrganizationMembers,
		Multiplex:   client.ByOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.OrganizationMember{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			{
				Name: "organization_id",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchOrganizationMembers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	includeAdmins := true
	result, _, err := c.Orgs.ListMembers(ctx, c.OrganizationID, includeAdmins)
	if err != nil {
		return err
	}

	for _, member := range result {
		res <- member
	}

	return nil
}
