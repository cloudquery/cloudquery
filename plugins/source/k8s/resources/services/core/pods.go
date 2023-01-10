package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Pods() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_pods",
		Resolver:  fetchPods,
		Multiplex: client.ContextMultiplex,
		Transform: transformers.TransformWithStruct(&v1.Pod{},
			client.SharedTransformersWithMoreSkipFields([]string{
				"DeprecatedServiceAccount", // Deprecated
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
				Name:     "status_host_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Status.HostIP"),
			},
			{
				Name:     "status_pod_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Status.PodIP"),
			},
			{
				Name:     "status_pod_ips",
				Type:     schema.TypeInetArray,
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
