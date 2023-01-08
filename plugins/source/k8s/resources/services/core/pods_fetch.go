package core

// func resolveCorePodPodIPs(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
// 	pod := resource.Item.(corev1.Pod)
// 	ips := make([]string, 0)

// 	for _, ip_struct := range pod.Status.PodIPs {
// 		ips = append(ips, ip_struct.IP)
// 	}

// 	return resource.Set(c.Name, ips)
// }
