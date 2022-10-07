package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fetchCoreNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	namespaces := meta.(*client.Client).Services().Namespaces
	opts := metav1.ListOptions{}
	for {
		result, err := namespaces.List(ctx, opts)
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
