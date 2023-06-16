package client

import (
	"context"
)

func (c *Client) initOrganizations(ctx context.Context) error {
	if len(c.Organizations) > 0 {
		// already set
		return nil
	}

	orgs, _, err := c.Orgs.List(ctx)
	if err != nil {
		return err
	}
	c.Organizations = orgs
	return nil
}

func (c *Client) WantOrganization(organizationID string) bool {
	for _, org := range c.Organizations {
		if org.ID == organizationID {
			return true
		}
	}
	return false
}
