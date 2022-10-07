package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fetchCorePods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	pods := meta.(*client.Client).Services().Pods
	result, err := pods.List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	res <- result.Items
	return nil
}

func resolveCorePodPodIPs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pod := resource.Item.(corev1.Pod)
	ips := make([]string, 0)

	for _, ip_struct := range pod.Status.PodIPs {
		ips = append(ips, ip_struct.IP)
	}

	return resource.Set(c.Name, ips)
}
