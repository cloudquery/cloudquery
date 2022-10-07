package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fetchCoreNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	nodes := meta.(*client.Client).Services().Nodes
	opts := metav1.ListOptions{}
	for {
		result, err := nodes.List(ctx, opts)
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
