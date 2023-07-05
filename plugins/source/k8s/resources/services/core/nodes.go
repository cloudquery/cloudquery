package core

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_nodes",
		Resolver:  fetchNodes,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.Node{},
			client.WithMoreSkipFields("DoNotUseExternalID"),
			transformers.WithPrimaryKeys("UID"),
		),
		Columns: schema.ColumnList{
			client.ContextColumn,
			{
				Name:     "spec_pod_cidr",
				Type:     types.ExtensionTypes.Inet,
				Resolver: client.StringToCidrPathResolver("Spec.PodCIDR"),
			},
			{
				Name:     "spec_pod_cidrs",
				Type:     arrow.ListOf(types.ExtensionTypes.Inet),
				Resolver: client.StringToCidrArrayPathResolver("Spec.PodCIDRs"),
			},
		},
	}
}

func fetchNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().CoreV1().Nodes()

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
