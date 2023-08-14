package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type TestClient struct {
	Logger   zerolog.Logger
	Spec     Spec
	ClientID int
}

func (*TestClient) ID() string {
	return "TestClient"
}

func (c *TestClient) withClientID(i int) *TestClient {
	t := *c
	t.ClientID = i
	return &t
}

func MultiplexBySpec(meta schema.ClientMeta) []schema.ClientMeta {
	cl := meta.(*TestClient)
	clients := make([]schema.ClientMeta, cl.Spec.NumClients)
	for i := 0; i < cl.Spec.NumClients; i++ {
		clients[i] = cl.withClientID(i)
	}
	return clients
}

func ResolveClientID(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*TestClient)
	return resource.Set(c.Name, cl.ClientID)
}
