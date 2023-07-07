package group

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name: "snyk_groups",
		Description: `https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to
	
This table lists all groups for the selected organizations. It uses the list organizations endpoint from the Snyk API.`,
		Resolver:  fetchGroups,
		Multiplex: client.SingleOrganization,
		Transform: transformers.TransformWithStruct(&snyk.Group{},
			transformers.WithPrimaryKeys("ID"),
		),
		Relations: []*schema.Table{
			groupMembers(),
		},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	groups := map[string]*snyk.Group{}
	for _, org := range c.Organizations {
		if org.Group != nil {
			groups[org.Group.ID] = org.Group
		}
	}
	for _, group := range groups {
		res <- group
	}
	return nil
}
