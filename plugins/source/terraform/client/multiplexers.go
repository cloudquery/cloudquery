package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

// BackendMultiplex add the ability to handle multiple backends in the same fetch
func BackendMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for name := range client.Backends {
		l = append(l, client.withSpecificBackend(name))
	}
	return l
}
