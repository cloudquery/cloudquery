package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func ObjectTypeMultiplex(objectTypes []string) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		cl := meta.(*Client)

		clients := make([]schema.ClientMeta, 0, len(objectTypes))

		for _, objectType := range objectTypes {
			clients = append(clients, cl.withObjectType(objectType))
		}

		return clients
	}
}
