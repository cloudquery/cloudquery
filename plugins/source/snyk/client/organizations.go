package client

import (
	"context"

	"golang.org/x/exp/slices"
)

func (c *Client) initOrganizations(ctx context.Context) error {
	if len(c.organizations) > 0 {
		// already set
		return nil
	}

	orgs, _, err := c.Orgs.List(ctx)
	if err != nil {
		return err
	}
	organizations := make([]string, 0, len(orgs))
	for _, org := range orgs {
		organizations = append(organizations, org.ID)
	}

	c.organizations = organizations
	return nil
}

func (c *Client) WantOrganization(organizationID string) bool {
	return slices.Contains(c.organizations, organizationID)
}
