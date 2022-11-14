package core

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	corev1 "k8s.io/api/core/v1"
)

func resolveCorePodPodIPs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pod := resource.Item.(corev1.Pod)
	ips := make([]string, 0)

	for _, ip_struct := range pod.Status.PodIPs {
		ips = append(ips, ip_struct.IP)
	}

	return resource.Set(c.Name, ips)
}
