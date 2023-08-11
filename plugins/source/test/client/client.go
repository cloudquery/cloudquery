package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger   zerolog.Logger
	Spec     Spec
	ClientID int
}

func (c *Client) ID() string {
	return fmt.Sprintf("Client#%d", c.ClientID+1)
}

func (c *Client) withClientID(i int) *Client {
	t := *c
	t.ClientID = i
	return &t
}

func MultiplexBySpec(meta schema.ClientMeta) []schema.ClientMeta {
	cl := meta.(*Client)
	clients := make([]schema.ClientMeta, cl.Spec.NumClients)
	for i := 0; i < cl.Spec.NumClients; i++ {
		clients[i] = cl.withClientID(i + 1)
	}
	return clients
}

func ResolveClientID(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*Client)
	return resource.Set(c.Name, cl.ClientID)
}
