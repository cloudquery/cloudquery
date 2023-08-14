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

func Pods() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_pods",
		Resolver:  fetchPods,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.Pod{},
			client.WithMoreSkipFields("DeprecatedServiceAccount"),
			transformers.WithPrimaryKeys("UID"),
		),
		Columns: schema.ColumnList{
			client.ContextColumn,
			{
				Name:     "status_host_ip",
				Type:     types.ExtensionTypes.Inet,
				Resolver: client.StringToInetPathResolver("Status.HostIP"),
			},
			{
				Name:     "status_pod_ip",
				Type:     types.ExtensionTypes.Inet,
				Resolver: client.StringToInetPathResolver("Status.PodIP"),
			},
			{
				Name:     "status_pod_ips",
				Type:     arrow.ListOf(types.ExtensionTypes.Inet),
				Resolver: resolveCorePodPodIPs,
			},
		},
	}
}

func fetchPods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().CoreV1().Pods("")

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

func resolveCorePodPodIPs(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pod := resource.Item.(v1.Pod)
	ips := make([]string, 0)

	for _, ip_struct := range pod.Status.PodIPs {
		ips = append(ips, ip_struct.IP)
	}

	return resource.Set(c.Name, ips)
}
