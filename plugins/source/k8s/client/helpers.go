package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type OwnerReferences struct {
	ResourceUID types.UID
	v1.OwnerReference
}

// ContextMultiplex returns a list of clients for each context from the cq config
func ContextMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	clients := make([]schema.ClientMeta, 0)
	for _, ctxName := range client.contexts {
		clients = append(clients, client.WithContext(ctxName))
	}
	return clients
}

// APIFilterContextMultiplex returns a list of clients for each context from the cq config
func APIFilterContextMultiplex(path string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)

		// in kubernetes version below 1.4 paths is nil
		if client.paths != nil {
			if _, ok := client.paths[path]; !ok {
				client.Logger().Warn().Str("path", path).Msg("The resource is not supported by current version of k8s")
				return []schema.ClientMeta{}
			}
		}

		clients := make([]schema.ClientMeta, 0, len(client.contexts))
		for _, ctxName := range client.contexts {
			clients = append(clients, client.WithContext(ctxName))
		}
		return clients
	}
}

// ResolveContext is a resolver that fills the k8s context field.
func ResolveContext(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Context)
}

// In k8s, IP Addresses may sometimes be empty-strings - but postgresql doesn't like that.
// So, the resolver for ip-addresses should recognize that case and not set null instead.
func StringToInetPathResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		value := funk.Get(r.Item, path, funk.WithAllowZero())

		stringValue, ok := value.(string)
		if ok && stringValue != "" {
			return r.Set(c.Name, value)
		}
		return r.Set(c.Name, nil)
	}
}

// In k8s, IP Addresses may sometimes be empty-strings - but postgresql doesn't like that.
// So, the resolver for ip-addresses should recognize that case and not set null instead.
func StringToCidrPathResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		value := funk.Get(r.Item, path, funk.WithAllowZero())

		stringValue, ok := value.(string)
		if ok && stringValue != "" {
			return r.Set(c.Name, value)
		}
		return r.Set(c.Name, nil)
	}
}
