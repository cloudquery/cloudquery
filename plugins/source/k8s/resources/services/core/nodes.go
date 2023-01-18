package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_nodes",
		Resolver:  fetchNodes,
		Multiplex: client.ContextMultiplex,
		Transform: transformers.TransformWithStruct(&v1.Node{},
			client.SharedTransformersWithMoreSkipFields([]string{
				"DoNotUseExternalID", // Deprecated
			})...),
		Columns: []schema.Column{
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: client.ResolveContext,
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "spec_pod_cidr",
				Type:     schema.TypeCIDR,
				Resolver: client.StringToCidrPathResolver("Spec.PodCIDR"),
			},
			{
				Name:     "spec_pod_cidrs",
				Type:     schema.TypeCIDRArray,
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
