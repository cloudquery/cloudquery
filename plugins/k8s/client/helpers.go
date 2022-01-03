package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/thoas/go-funk"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	ContextFieldName = "context"
	ContextFieldDesc = "Name of the context from k8s configuration."
)

var CommonContextField = schema.Column{
	Name:        ContextFieldName,
	Description: ContextFieldDesc,
	Type:        schema.TypeString,
	Resolver:    ResolveContext,
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

// DeleteContextFilter returns a delete filter that cleans up the data belonging to the k8s context.
func DeleteContextFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{ContextFieldName, client.Context}
}

// ResolveContext is a resolver that fills the k8s context field.
func ResolveContext(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Context)
}

func OwnerReferenceResolver(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	v := funk.Get(parent.Item, "ObjectMeta")
	if v == nil {
		return nil
	}
	objMeta := v.(v1.ObjectMeta)
	if len(objMeta.OwnerReferences) == 0 {
		return nil
	}
	refs := make([]OwnerReferences, len(objMeta.OwnerReferences))
	for i, o := range objMeta.OwnerReferences {
		refs[i] = OwnerReferences{
			ResourceUID:    parent.Get("uid").(types.UID),
			OwnerReference: o,
		}
	}
	res <- refs
	return nil
}

type OwnerReferences struct {
	ResourceUID types.UID
	v1.OwnerReference
}
