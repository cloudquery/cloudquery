package networking

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NetworkPolicies() *schema.Table {
	return &schema.Table{
		Name:         "k8s_networking_network_policies",
		Description:  "NetworkPolicy describes what network traffic is allowed for a set of Pods",
		Resolver:     fetchNetworkingNetworkPolicies,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbiddenNotFound,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.APIVersion"),
			},
			{
				Name:        "name",
				Description: "Name must be unique within a namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:        "generate_name",
				Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.GenerateName"),
			},
			{
				Name:        "namespace",
				Description: "Namespace defines the space within which each name must be unique",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:        "self_link",
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.SelfLink"),
			},
			{
				Name:        "uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:        "resource_version",
				Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "generation",
				Description: "A sequence number representing a specific generation of the desired state. Populated by the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:          "deletion_grace_period_seconds",
				Description:   "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
				IgnoreInTests: true,
			},
			{
				Name:          "labels",
				Description:   "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ObjectMeta.Labels"),
				IgnoreInTests: true,
			},
			{
				Name:          "annotations",
				Description:   "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ObjectMeta.Annotations"),
				IgnoreInTests: true,
			},
			{
				Name:          "owner_references",
				Description:   "List of objects depended by this object",
				Type:          schema.TypeJSON,
				Resolver:      resolveNetworkingNetworkPoliciesOwnerReferences,
				IgnoreInTests: true,
			},
			{
				Name:          "finalizers",
				Description:   "Must be empty before the object is deleted from the registry",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ObjectMeta.Finalizers"),
				IgnoreInTests: true,
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "managed_fields",
				Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkingNetworkPoliciesManagedFields,
			},
			{
				Name:          "pod_selector_match_labels",
				Description:   "matchLabels is a map of {key,value} pairs",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Spec.PodSelector.MatchLabels"),
				IgnoreInTests: true,
			},
			{
				Name:        "policy_types",
				Description: "List of rule types that the NetworkPolicy relates to. Valid options are [\"Ingress\"], [\"Egress\"], or [\"Ingress\", \"Egress\"]. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ \"Egress\" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include \"Egress\" (since such a policy would not include an Egress section and would otherwise default to just [ \"Ingress\" ]). This field is beta-level in 1.8 +optional",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.PolicyTypes"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_networking_network_policy_pod_selector_match_expressions",
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:    fetchNetworkingNetworkPolicyPodSelectorMatchExpressions,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "values is an array of string values",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "k8s_networking_network_policy_ingress",
				Description: "NetworkPolicyIngressRule describes a particular set of traffic that is allowed to the pods matched by a NetworkPolicySpec's podSelector",
				Resolver:    fetchNetworkingNetworkPolicyIngresses,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "network_policy_uid",
						Description: "Unique internal ID of Network Policy resource",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("uid"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_networking_network_policy_ingress_ports",
						Description: "NetworkPolicyPort describes a port to allow traffic on",
						Resolver:    fetchNetworkingNetworkPolicyIngressPorts,
						Columns: []schema.Column{
							{
								Name:        "network_policy_ingress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_ingress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "protocol",
								Description: "The protocol (TCP, UDP, or SCTP) which traffic must match",
								Type:        schema.TypeString,
							},
							{
								Name:     "port_type",
								Type:     schema.TypeBigInt,
								Resolver: schema.PathResolver("Port.Type"),
							},
							{
								Name:     "port_int_val",
								Type:     schema.TypeInt,
								Resolver: schema.PathResolver("Port.IntVal"),
							},
							{
								Name:     "port_str_val",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Port.StrVal"),
							},
							{
								Name:          "end_port",
								Description:   "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy",
								Type:          schema.TypeInt,
								IgnoreInTests: true,
							},
						},
					},
					{
						Name:        "k8s_networking_network_policy_ingress_from",
						Description: "NetworkPolicyPeer describes a peer to allow traffic to/from",
						Resolver:    fetchNetworkingNetworkPolicyIngressFroms,
						Columns: []schema.Column{
							{
								Name:        "network_policy_ingress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_ingress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:          "pod_selector_match_labels",
								Description:   "matchLabels is a map of {key,value} pairs",
								Type:          schema.TypeJSON,
								Resolver:      schema.PathResolver("PodSelector.MatchLabels"),
								IgnoreInTests: true,
							},
							{
								Name:          "pod_selector_match_expressions",
								Description:   "matchExpressions is a list of label selector requirements",
								Type:          schema.TypeJSON,
								Resolver:      resolveNetworkingNetworkPolicyIngressFromsPodSelectorMatchExpressions,
								IgnoreInTests: true,
							},
							{
								Name:        "namespace_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("NamespaceSelector.MatchLabels"),
							},
							{
								Name:          "namespace_selector_match_expressions",
								Description:   "matchExpressions is a list of label selector requirements",
								Type:          schema.TypeJSON,
								Resolver:      resolveNetworkingNetworkPolicyIngressFromsNamespaceSelectorMatchExpressions,
								IgnoreInTests: true,
							},
							{
								Name:        "ip_block_cidr",
								Description: "CIDR is a string representing the IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("IPBlock.CIDR"),
							},
							{
								Name:        "ip_block_except",
								Description: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\" Except values will be rejected if they are outside the CIDR range +optional",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("IPBlock.Except"),
							},
						},
					},
				},
			},
			{
				Name:        "k8s_networking_network_policy_egress",
				Description: "NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched by a NetworkPolicySpec's podSelector",
				Resolver:    fetchNetworkingNetworkPolicyEgresses,
				Columns: []schema.Column{
					{
						Name:        "network_policy_cq_id",
						Description: "Unique CloudQuery ID of k8s_networking_network_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "network_policy_uid",
						Description: "Unique internal ID of Network Policy resource",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("uid"),
					},
				},
				Relations: []*schema.Table{
					{
						IgnoreInTests: true,
						Name:          "k8s_networking_network_policy_egress_ports",
						Description:   "NetworkPolicyPort describes a port to allow traffic on",
						Resolver:      fetchNetworkingNetworkPolicyEgressPorts,
						Columns: []schema.Column{
							{
								Name:        "network_policy_egress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_egress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "protocol",
								Description: "The protocol (TCP, UDP, or SCTP) which traffic must match",
								Type:        schema.TypeString,
							},
							{
								Name:     "port_type",
								Type:     schema.TypeBigInt,
								Resolver: schema.PathResolver("Port.Type"),
							},
							{
								Name:     "port_int_val",
								Type:     schema.TypeInt,
								Resolver: schema.PathResolver("Port.IntVal"),
							},
							{
								Name:     "port_str_val",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Port.StrVal"),
							},
							{
								Name:        "end_port",
								Description: "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy",
								Type:        schema.TypeInt,
							},
						},
					},
					{
						IgnoreInTests: true,
						Name:          "k8s_networking_network_policy_egress_to",
						Description:   "NetworkPolicyPeer describes a peer to allow traffic to/from",
						Resolver:      fetchNetworkingNetworkPolicyEgressTos,
						Columns: []schema.Column{
							{
								Name:        "network_policy_egress_cq_id",
								Description: "Unique CloudQuery ID of k8s_networking_network_policy_egress table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "pod_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("PodSelector.MatchLabels"),
							},
							{
								Name:        "pod_selector_match_expressions",
								Description: "matchExpressions is a list of label selector requirements",
								Type:        schema.TypeJSON,
								Resolver:    resolveNetworkingNetworkPolicyEgressTosPodSelectorMatchExpressions,
							},
							{
								Name:        "namespace_selector_match_labels",
								Description: "matchLabels is a map of {key,value} pairs",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("NamespaceSelector.MatchLabels"),
							},
							{
								Name:        "namespace_selector_match_expressions",
								Description: "matchExpressions is a list of label selector requirements",
								Type:        schema.TypeJSON,
								Resolver:    resolveNetworkingNetworkPolicyEgressTosNamespaceSelectorMatchExpressions,
							},
							{
								Name:        "ip_block_cidr",
								Description: "CIDR is a string representing the IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\"",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("IPBlock.CIDR"),
							},
							{
								Name:        "ip_block_except",
								Description: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\" Except values will be rejected if they are outside the CIDR range +optional",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("IPBlock.Except"),
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkingNetworkPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client).Services().NetworkPolicies
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
func resolveNetworkingNetworkPoliciesOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveNetworkingNetworkPoliciesManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchNetworkingNetworkPolicyPodSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.Spec.PodSelector.MatchExpressions
	return nil
}
func fetchNetworkingNetworkPolicyIngresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.Spec.Ingress
	return nil
}
func fetchNetworkingNetworkPolicyIngressPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyIngressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchNetworkingNetworkPolicyIngressFroms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyIngressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.From
	return nil
}
func resolveNetworkingNetworkPolicyIngressFromsPodSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", resource.Item)
	}
	if p.PodSelector == nil {
		return nil
	}
	b, err := json.Marshal(p.PodSelector.MatchExpressions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveNetworkingNetworkPolicyIngressFromsNamespaceSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", resource.Item)
	}
	if p.NamespaceSelector == nil {
		return nil
	}
	b, err := json.Marshal(p.NamespaceSelector.MatchExpressions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchNetworkingNetworkPolicyEgresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicy instance: %T", parent.Item)
	}
	res <- p.Spec.Egress
	return nil
}
func fetchNetworkingNetworkPolicyEgressPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyEgressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.Ports
	return nil
}
func fetchNetworkingNetworkPolicyEgressTos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(networkingv1.NetworkPolicyEgressRule)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyIngressRule instance: %T", parent.Item)
	}
	res <- p.To
	return nil
}
func resolveNetworkingNetworkPolicyEgressTosPodSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", resource.Item)
	}
	if p.PodSelector == nil {
		return nil
	}
	b, err := json.Marshal(p.PodSelector.MatchExpressions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveNetworkingNetworkPolicyEgressTosNamespaceSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(networkingv1.NetworkPolicyPeer)
	if !ok {
		return fmt.Errorf("not a networkingv1.NetworkPolicyPeer instance: %T", resource.Item)
	}
	if p.NamespaceSelector == nil {
		return nil
	}
	b, err := json.Marshal(p.NamespaceSelector.MatchExpressions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
