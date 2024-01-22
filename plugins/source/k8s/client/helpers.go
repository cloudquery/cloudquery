package client

import (
	"context"
	"reflect"

	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func ContextNamespaceMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	clients := make([]schema.ClientMeta, 0)
	for _, ctxName := range client.contexts {
		for _, ns := range client.namespaces[ctxName] {
			clients = append(clients, client.WithContext(ctxName).WithNamespace(ns.Name))
		}
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

// StringToNullablePathResolver is used to set values of "" & "None" to `nil`.
func StringToNullablePathResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		value := funk.Get(r.Item, path, funk.WithAllowZero())

		value, ok := value.(string)
		if !ok {
			return r.Set(c.Name, nil)
		}

		switch value {
		case "", "None":
			return r.Set(c.Name, nil)
		default:
			return r.Set(c.Name, value)
		}
	}
}

// StringToNullableArrayPathResolver is the same as StringToNullablePathResolver but for slices.
// The result will be `[]*string`
func StringToNullableArrayPathResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		value := funk.Get(r.Item, path, funk.WithAllowZero())

		stringArrayValue, ok := value.([]string)
		if !ok {
			return r.Set(c.Name, nil)
		}

		sanitized := make([]*string, len(stringArrayValue))

		for i := range stringArrayValue {
			switch stringArrayValue[i] {
			case "", "None":
			// nop, as already nil
			default:
				sanitized[i] = &stringArrayValue[i]
			}
		}

		return r.Set(c.Name, sanitized)
	}
}

// isK8sTimeStruct returns true if the given type is a metav1.Time struct or a pointer to it.
func isK8sTimeStruct(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Ptr {
		return isK8sTimeStruct(fieldType.Elem())
	}

	if fieldKind == reflect.Struct && fieldType == reflect.TypeOf(v1.Time{}) {
		return true
	}

	return false
}
