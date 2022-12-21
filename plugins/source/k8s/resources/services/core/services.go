// Code generated by codegen; DO NOT EDIT.

package core

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_services",
		Resolver:  fetchServices,
		Multiplex: client.ContextMultiplex,
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
				Resolver: schema.PathResolver("Spec.ClusterIPs"),
			},
			{
				Name:     "spec_external_ips",
				Type:     schema.TypeInetArray,
				Resolver: schema.PathResolver("Spec.ExternalIPs"),
			},
			{
				Name:     "spec_load_balancer_ip",
				Type:     schema.TypeInet,
				Resolver: client.StringToInetPathResolver("Spec.LoadBalancerIP"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "owner_references",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OwnerReferences"),
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Finalizers"),
			},
			{
				Name:     "spec_ports",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Ports"),
			},
			{
				Name:     "spec_selector",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Selector"),
			},
			{
				Name:     "spec_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.Type"),
			},
			{
				Name:     "spec_session_affinity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.SessionAffinity"),
			},
			{
				Name:     "spec_load_balancer_source_ranges",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.LoadBalancerSourceRanges"),
			},
			{
				Name:     "spec_external_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ExternalName"),
			},
			{
				Name:     "spec_external_traffic_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ExternalTrafficPolicy"),
			},
			{
				Name:     "spec_health_check_node_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.HealthCheckNodePort"),
			},
			{
				Name:     "spec_publish_not_ready_addresses",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.PublishNotReadyAddresses"),
			},
			{
				Name:     "spec_session_affinity_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.SessionAffinityConfig"),
			},
			{
				Name:     "spec_ip_families",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.IPFamilies"),
			},
			{
				Name:     "spec_ip_family_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.IPFamilyPolicy"),
			},
			{
				Name:     "spec_allocate_load_balancer_node_ports",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.AllocateLoadBalancerNodePorts"),
			},
			{
				Name:     "spec_load_balancer_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.LoadBalancerClass"),
			},
			{
				Name:     "spec_internal_traffic_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.InternalTrafficPolicy"),
			},
			{
				Name:     "status_load_balancer",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.LoadBalancer"),
			},
			{
				Name:     "status_conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Conditions"),
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
