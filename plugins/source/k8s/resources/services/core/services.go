package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_services",
		Resolver:  fetchServices,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.Service{}),
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
				Name:     "spec_cluster_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Spec.ClusterIP"),
			},
			{
				Name:     "spec_cluster_ips",
				Type:     schema.TypeInetArray,
				Resolver: client.StringToInetArrayPathResolver("Spec.ClusterIPs"),
			},
			{
				Name:     "spec_external_ips",
				Type:     schema.TypeInetArray,
				Resolver: client.StringToInetArrayPathResolver("Spec.ExternalIPs"),
			},
			{
				Name:     "spec_load_balancer_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Spec.LoadBalancerIP"),
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().CoreV1().Services("")

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
