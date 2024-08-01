package nodes

import (
	"context"

	v1 "k8s.io/api/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RuntimeClasses() *schema.Table {
	return &schema.Table{
		Name:      "k8s_nodes_runtime_classes",
		Resolver:  fetchRuntimeClasses,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.RuntimeClass{}, transformers.WithPrimaryKeys("UID")),
		Columns:   schema.ColumnList{client.ContextColumn},
	}
}

func fetchRuntimeClasses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).CoreAPI().NodeV1().RuntimeClasses()

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
