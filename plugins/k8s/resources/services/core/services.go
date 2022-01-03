package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_services",
		Description:  "Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.",
		Resolver:     fetchCoreServices,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "Defines the versioned schema of this representation of an object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.APIVersion"),
			},
			{
				Name:        "name",
				Description: "Unique name within a namespace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:        "namespace",
				Description: "Namespace defines the space within which each name must be unique.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:        "uid",
				Description: "UID is the unique in time and space value for this object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:        "resource_version",
				Description: "An opaque value that represents the internal version of this object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "generation",
				Description: "A sequence number representing a specific generation of the desired state.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:          "deletion_grace_period_seconds",
				Description:   "Number of seconds allowed for this object to gracefully terminate.",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
				IgnoreInTests: true,
			},
			{
				Name:        "labels",
				Description: "Map of string keys and values that can be used to organize and categorize objects.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:        "annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:          "owner_references",
				Description:   "List of objects depended by this object.",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreServiceOwnerReferences,
				IgnoreInTests: true,
			},
			{
				Name:          "finalizers",
				Description:   "List of finalizers",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ObjectMeta.Finalizers"),
				IgnoreInTests: true,
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "selector",
				Description: "Route service traffic to pods with label keys and values matching this selector",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Selector"),
			},
			{
				Name:        "cluster_ip",
				Description: "clusterIP is the IP address of the service and is usually assigned randomly",
				Type:        schema.TypeInet,
				Resolver:    resolveCoreServicesClusterIP,
			},
			{
				Name:        "cluster_ips",
				Description: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly",
				Type:        schema.TypeInetArray,
				Resolver:    resolveCoreServicesClusterIPs,
			},
			{
				Name:        "type",
				Description: "type determines how the Service is exposed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Type"),
			},
			{
				Name:        "external_ips",
				Description: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service",
				Type:        schema.TypeInetArray,
				Resolver:    resolveCoreServicesExternalIPs,
			},
			{
				Name:        "session_affinity",
				Description: "Used to maintain session affinity.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.SessionAffinity"),
			},
			{
				Name:        "load_balancer_ip",
				Description: "Load balancer will get created with the IP specified in this field.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.LoadBalancerIP"),
			},
			{
				Name:          "load_balancer_source_ranges",
				Description:   "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer to the specified client IPs",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Spec.LoadBalancerSourceRanges"),
				IgnoreInTests: true,
			},
			{
				Name:        "external_name",
				Description: "The external reference that discovery mechanisms will return as an alias for this service.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ExternalName"),
			},
			{
				Name:        "external_traffic_policy",
				Description: "Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ExternalTrafficPolicy"),
			},
			{
				Name:        "health_check_node_port",
				Description: "Specifies the healthcheck nodePort for the service.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.HealthCheckNodePort"),
			},
			{
				Name:        "publish_not_ready_addresses",
				Description: "Indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.PublishNotReadyAddresses"),
			},
			{
				Name:        "session_affinity_config_client_ip_timeout_seconds",
				Description: "Specifies the seconds of ClientIP type session sticky time.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds"),
			},
			{
				Name:        "ip_families",
				Description: "A list of IP families assigned to this service.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.IPFamilies"),
			},
			{
				Name:        "ip_family_policy",
				Description: "Represents the dual-stack-ness requested or required by this Service.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.IPFamilyPolicy"),
			},
			{
				Name:          "allocate_load_balancer_node_ports",
				Description:   "Defines if NodePorts will be automatically allocated for services with type LoadBalancer",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("Spec.AllocateLoadBalancerNodePorts"),
				IgnoreInTests: true,
			},
			{
				Name:          "load_balancer_class",
				Description:   "The class of the load balancer implementation this Service belongs to.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Spec.LoadBalancerClass"),
				IgnoreInTests: true,
			},
			{
				Name:        "internal_traffic_policy",
				Description: "Specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. \"Cluster\" routes internal traffic to a Service to all endpoints. \"Local\" routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.InternalTrafficPolicy"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_core_service_ports",
				Description: "The list of ports that are exposed by this service.",
				Resolver:    fetchCoreServicePorts,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of this port within the service",
						Type:        schema.TypeString,
					},
					{
						Name:        "protocol",
						Description: "The IP protocol for this port",
						Type:        schema.TypeString,
					},
					{
						Name:          "app_protocol",
						Description:   "The application protocol for this port.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "port",
						Description: "The port that will be exposed by this service.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "target_port_type",
						Description: "Port type, integer or string",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("TargetPort.Type"),
					},
					{
						Name:        "target_port_int_val",
						Description: "Port as an integer value.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("TargetPort.IntVal"),
					},
					{
						Name:        "target_port_str_val",
						Description: "Port as a string value.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetPort.StrVal"),
					},
					{
						Name:        "node_port",
						Description: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				IgnoreInTests: true,
				Name:          "k8s_core_service_load_balancer_ingresses",
				Description:   "LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.",
				Resolver:      fetchCoreServiceLoadBalancerIngresses,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ip",
						Description: "IP is set for load-balancer ingress points that are IP based.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IP"),
					},
					{
						Name:        "hostname",
						Description: "A set for load-balancer ingress points that are DNS based",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						IgnoreInTests: true,
						Name:          "k8s_core_service_load_balancer_ingress_ports",
						Resolver:      fetchCoreServiceLoadBalancerIngressPorts,
						Columns: []schema.Column{
							{
								Name:        "service_load_balancer_ingress_cq_id",
								Description: "Unique CloudQuery ID of k8s_core_service_load_balancer_ingresses table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "port",
								Description: "Port is the port number of the service port of which status is recorded here",
								Type:        schema.TypeInt,
							},
							{
								Name:        "protocol",
								Description: "Protocol is the protocol of the service port of which status is recorded here.",
								Type:        schema.TypeString,
							},
							{
								Name:        "error",
								Description: "Error is to record the problem with the service port.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				IgnoreInTests: true,
				Name:          "k8s_core_service_conditions",
				Description:   "Condition contains details for one aspect of the current state of this API Resource. --- This struct is intended for direct use as an array at the field path .status.conditions",
				Resolver:      fetchCoreServiceConditions,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "type of condition in CamelCase or in foo.example.com/CamelCase. --- Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt) +required +kubebuilder:validation:Required +kubebuilder:validation:Pattern=`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$` +kubebuilder:validation:MaxLength=316",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "status of the condition, one of True, False, Unknown. +required +kubebuilder:validation:Required +kubebuilder:validation:Enum=True;False;Unknown",
						Type:        schema.TypeString,
					},
					{
						Name:        "observed_generation",
						Description: "observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance. +optional +kubebuilder:validation:Minimum=0",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "reason",
						Description: "reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty. +required +kubebuilder:validation:Required +kubebuilder:validation:MaxLength=1024 +kubebuilder:validation:MinLength=1 +kubebuilder:validation:Pattern=`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "message is a human readable message indicating details about the transition. This may be an empty string. +required +kubebuilder:validation:Required +kubebuilder:validation:MaxLength=32768",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCoreServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	services := meta.(*client.Client).Services().Services
	opts := metav1.ListOptions{}
	for {
		result, err := services.List(ctx, metav1.ListOptions{})
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

func resolveCoreServicesClusterIP(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", resource.Item)
	}
	ip := net.ParseIP(service.Spec.ClusterIP)
	if ip != nil {
		if v4 := ip.To4(); v4 != nil {
			ip = v4
		}
	}
	return resource.Set(c.Name, ip)
}

func resolveCoreServicesClusterIPs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", resource.Item)
	}
	ips := make([]net.IP, 0, len(service.Spec.ClusterIPs))
	for _, v := range service.Spec.ClusterIPs {
		ip := net.ParseIP(v)
		if ip != nil {
			if v4 := ip.To4(); v4 != nil {
				ip = v4
			}
		}
		ips = append(ips, ip)
	}
	return resource.Set(c.Name, ips)
}

func resolveCoreServicesExternalIPs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", resource.Item)
	}
	ips := make([]net.IP, 0, len(service.Spec.ExternalIPs))
	for _, v := range service.Spec.ExternalIPs {
		ip := net.ParseIP(v)
		if ip != nil {
			if v4 := ip.To4(); v4 != nil {
				ip = v4
			}
		}
		ips = append(ips, ip)
	}
	return resource.Set(c.Name, ips)
}

func resolveCoreServiceOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", resource.Item)
	}
	b, err := json.Marshal(service.ObjectMeta.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchCoreServicePorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", parent.Item)
	}
	res <- service.Spec.Ports
	return nil
}

func fetchCoreServiceLoadBalancerIngresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", parent.Item)
	}
	res <- service.Status.LoadBalancer.Ingress
	return nil
}

func fetchCoreServiceLoadBalancerIngressPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ingress, ok := parent.Item.(corev1.LoadBalancerIngress)
	if !ok {
		return fmt.Errorf("not a corev1.LoadBalancerIngress instance: %T", parent.Item)
	}
	res <- ingress.Ports
	return nil
}

func fetchCoreServiceConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(corev1.Service)
	if !ok {
		return fmt.Errorf("not a corev1.Service instance: %T", parent.Item)
	}
	res <- service.Status.Conditions
	return nil
}
