package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_namespaces",
		Resolver:  fetchNamespaces,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.Namespace{}, transformers.WithPrimaryKeys("UID")),
		Columns:   schema.ColumnList{client.ContextColumn},
	}
}

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	if c.Namespaces() != nil {
		// If we already got the namespaces we can exit early
		res <- c.Namespaces()
		return nil
	}
	cl := c.Client().CoreV1().Namespaces()

	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
