package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func SingleOrganization(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	return []schema.ClientMeta{client.WithOrganization(client.organizations[0])}
}

func ByOrganization(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	l := make([]schema.ClientMeta, 0, len(client.organizations))
	for _, o := range client.organizations {
		l = append(l, client.WithOrganization(o))
	}
	return l
}

func (c *Client) WithOrganization(organizationID string) schema.ClientMeta {
	return &Client{
		Client:         c.Client,
		OrganizationID: organizationID,
		organizations:  c.organizations,
		logger:         c.logger.With().Str("organization", organizationID).Logger(),
	}
}
