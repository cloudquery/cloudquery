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

func Endpoints() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_endpoints",
		Description:  "Endpoints is a collection of endpoints that implement the actual service",
		Resolver:     fetchCoreEndpoints,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbidden,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.",
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
				Name:        "labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:        "annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:          "owner_references",
				Description:   "List of objects depended by this object",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreEndpointsOwnerReferences,
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
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "managed_fields",
				Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreEndpointsManagedFields,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_core_endpoint_subsets",
				Description: "EndpointSubset is a group of addresses with a common set of ports",
				Resolver:    fetchCoreEndpointSubsets,
				Columns: []schema.Column{
					{
						Name:        "endpoint_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_endpoints table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "endpoint_name",
						Description: "Name must be unique within a namespace",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ObjectMeta.Name"),
					},
					{
						Name:        "endpoint_uid",
						Description: "UID is the unique in time and space value for this object",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ObjectMeta.UID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_core_endpoint_subset_addresses",
						Description: "EndpointAddress is a tuple that describes single IP address.",
						Resolver:    fetchCoreEndpointSubsetAddresses,
						Columns: []schema.Column{
							{
								Name:        "endpoint_subset_cq_id",
								Description: "Unique CloudQuery ID of k8s_core_endpoint_subsets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "ip",
								Description: "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms",
								Type:        schema.TypeInet,
								Resolver:    resolveCoreEndpointSubsetAddressesIP,
							},
							{
								Name:        "hostname",
								Description: "The Hostname of this endpoint",
								Type:        schema.TypeString,
							},
							{
								Name:        "node_name",
								Description: "Optional: Node hosting this endpoint",
								Type:        schema.TypeString,
							},
							{
								Name:        "target_ref_kind",
								Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.Kind"),
							},
							{
								Name:        "target_ref_namespace",
								Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.Namespace"),
							},
							{
								Name:        "target_ref_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.Name"),
							},
							{
								Name:        "target_ref_uid",
								Description: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.UID"),
							},
							{
								Name:        "target_ref_api_version",
								Description: "API version of the referent.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.APIVersion"),
							},
							{
								Name:        "target_ref_resource_version",
								Description: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.ResourceVersion"),
							},
							{
								Name:        "target_ref_field_path",
								Description: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: \"spec.containers{name}\" (where \"name\" refers to the name of the container that triggered the event) or if no container name is specified \"spec.containers[2]\" (container with index 2 in this pod)",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.FieldPath"),
							},
						},
					},
					{
						Name:          "k8s_core_endpoint_subset_not_ready_addresses",
						Description:   "EndpointAddress is a tuple that describes single IP address.",
						Resolver:      fetchCoreEndpointSubsetNotReadyAddresses,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "endpoint_subset_cq_id",
								Description: "Unique CloudQuery ID of k8s_core_endpoint_subsets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "ip",
								Description: "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms",
								Type:        schema.TypeInet,
								Resolver:    resolveCoreEndpointSubsetNotReadyAddressesIP,
							},
							{
								Name:        "hostname",
								Description: "The Hostname of this endpoint",
								Type:        schema.TypeString,
							},
							{
								Name:        "node_name",
								Description: "Optional: Node hosting this endpoint",
								Type:        schema.TypeString,
							},
							{
								Name:        "target_ref_kind",
								Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.Kind"),
							},
							{
								Name:        "target_ref_namespace",
								Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.Namespace"),
							},
							{
								Name:        "target_ref_name",
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.Name"),
							},
							{
								Name:        "target_ref_uid",
								Description: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.UID"),
							},
							{
								Name:        "target_ref_api_version",
								Description: "API version of the referent.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.APIVersion"),
							},
							{
								Name:        "target_ref_resource_version",
								Description: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.ResourceVersion"),
							},
							{
								Name:        "target_ref_field_path",
								Description: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: \"spec.containers{name}\" (where \"name\" refers to the name of the container that triggered the event) or if no container name is specified \"spec.containers[2]\" (container with index 2 in this pod)",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TargetRef.FieldPath"),
							},
						},
					},
					{
						Name:        "k8s_core_endpoint_subset_ports",
						Description: "EndpointPort is a tuple that describes a single port.",
						Resolver:    fetchCoreEndpointSubsetPorts,
						Columns: []schema.Column{
							{
								Name:        "endpoint_subset_cq_id",
								Description: "Unique CloudQuery ID of k8s_core_endpoint_subsets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The name of this port",
								Type:        schema.TypeString,
							},
							{
								Name:        "port",
								Description: "The port number of the endpoint.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "protocol",
								Description: "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
								Type:        schema.TypeString,
							},
							{
								Name:          "app_protocol",
								Description:   "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
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

func fetchCoreEndpoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client).Services().Endpoints
	opts := metav1.ListOptions{}
	for {
		result, err := c.List(ctx, opts)
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

func resolveCoreEndpointsOwnerReferences(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Endpoints)
	if !ok {
		return fmt.Errorf("not a corev1.Endpoints instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreEndpointsManagedFields(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(corev1.Endpoints)
	if !ok {
		return fmt.Errorf("not a corev1.Endpoints instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreEndpointSubsetAddressesIP(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	endpointAddress, ok := resource.Item.(corev1.EndpointAddress)
	if !ok {
		return fmt.Errorf("not a corev1.EndpointAddress instance: %T", resource.Item)
	}
	ip := net.ParseIP(endpointAddress.IP)
	if ip != nil {
		if v4 := ip.To4(); v4 != nil {
			ip = v4
		}
	}
	return resource.Set(c.Name, ip)
}

func resolveCoreEndpointSubsetNotReadyAddressesIP(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	endpointAddress, ok := resource.Item.(corev1.EndpointAddress)
	if !ok {
		return fmt.Errorf("not a corev1.EndpointAddress instance: %T", resource.Item)
	}
	ip := net.ParseIP(endpointAddress.IP)
	if ip != nil {
		if v4 := ip.To4(); v4 != nil {
			ip = v4
		}
	}
	return resource.Set(c.Name, ip)
}

func fetchCoreEndpointSubsets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	endpoints, ok := parent.Item.(corev1.Endpoints)
	if !ok {
		return fmt.Errorf("not a corev1.Endpoints instance: %T", parent.Item)
	}
	res <- endpoints.Subsets
	return nil
}

func fetchCoreEndpointSubsetAddresses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	endpointSubset, ok := parent.Item.(corev1.EndpointSubset)
	if !ok {
		return fmt.Errorf("not a corev1.EndpointSubset instance: %T", parent.Item)
	}
	res <- endpointSubset.Addresses
	return nil
}

func fetchCoreEndpointSubsetNotReadyAddresses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	endpointSubset, ok := parent.Item.(corev1.EndpointSubset)
	if !ok {
		return fmt.Errorf("not a corev1.EndpointSubset instance: %T", parent.Item)
	}
	res <- endpointSubset.NotReadyAddresses
	return nil
}

func fetchCoreEndpointSubsetPorts(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	endpointSubset, ok := parent.Item.(corev1.EndpointSubset)
	if !ok {
		return fmt.Errorf("not a corev1.EndpointSubset instance: %T", parent.Item)
	}
	res <- endpointSubset.Ports
	return nil
}
