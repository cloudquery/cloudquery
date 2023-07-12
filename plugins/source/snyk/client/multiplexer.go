package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func SingleOrganization(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	return []schema.ClientMeta{client.WithOrganization(client.Organizations[0].ID)}
}

func ByOrganization(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	l := make([]schema.ClientMeta, 0, len(client.Organizations))
	for _, o := range client.Organizations {
		l = append(l, client.WithOrganization(o.ID))
	}
	return l
}

func (c *Client) WithOrganization(organizationID string) schema.ClientMeta {
	return &Client{
		Client:         c.Client,
		Spec:           c.Spec,
		OrganizationID: organizationID,
		Organizations:  c.Organizations,
		logger:         c.logger.With().Str("organization", organizationID).Logger(),
		maxRetries:     c.maxRetries,
		backoff:        c.backoff,
	}
}
