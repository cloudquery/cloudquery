package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:         "k8s_core_nodes",
		Description:  "Node is a worker node in Kubernetes.",
		Resolver:     fetchCoreNodes,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		IgnoreError:  client.IgnoreForbiddenNotFound,
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
				Resolver:      resolveCoreNodeOwnerReferences,
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
				Name:        "pod_cidr",
				Description: "Represents the pod IP range assigned to the node.",
				Type:        schema.TypeCIDR,
				Resolver:    resolveCoreNodePodCIDR,
			},
			{
				Name:        "pod_cidrs",
				Description: "Represents the IP ranges assigned to the node for usage by Pods on that node",
				Type:        schema.TypeCIDRArray,
				Resolver:    resolveCoreNodePodCIDRs,
			},
			{
				Name:        "provider_id",
				Description: "ID of the node assigned by the cloud provider.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ProviderID"),
			},
			{
				Name:        "unschedulable",
				Description: "Unschedulable controls node schedulability of new pods",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Unschedulable"),
			},
			{
				Name:          "taints",
				Description:   "If specified, the node's taints.",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreNodeTaints,
				IgnoreInTests: true,
			},
			{
				Name:        "capacity",
				Description: "Capacity represents the total resources of a node.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Status.Capacity"),
			},
			{
				Name:        "allocatable",
				Description: "Allocatable represents the resources of a node that are available for scheduling.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Status.Allocatable"),
			},
			{
				Name:        "phase",
				Description: "NodePhase is the recently observed lifecycle phase of the node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Phase"),
			},
			{
				Name:        "conditions",
				Description: "Conditions is an array of current observed node conditions.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreNodeConditions,
			},
			{
				Name:        "daemon_endpoints_kubelet_endpoint_port",
				Description: "Port number of the given endpoint.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.DaemonEndpoints.KubeletEndpoint.Port"),
			},
			{
				Name:        "machine_id",
				Description: "MachineID reported by the node",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.MachineID"),
			},
			{
				Name:        "system_uuid",
				Description: "SystemUUID reported by the node",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.SystemUUID"),
			},
			{
				Name:        "boot_id",
				Description: "Boot ID reported by the node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.BootID"),
			},
			{
				Name:        "kernel_version",
				Description: "Kernel Version reported by the node from 'uname -r'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.KernelVersion"),
			},
			{
				Name:        "os_image",
				Description: "OS Image reported by the node from /etc/os-release",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.OSImage"),
			},
			{
				Name:        "container_runtime_version",
				Description: "Container runtime version reported by the node through runtime remote API.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.ContainerRuntimeVersion"),
			},
			{
				Name:        "kubelet_version",
				Description: "Kubelet Version reported by the node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.KubeletVersion"),
			},
			{
				Name:        "kube_proxy_version",
				Description: "KubeProxy Version reported by the node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.KubeProxyVersion"),
			},
			{
				Name:        "operating_system",
				Description: "The Operating System reported by the node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.OperatingSystem"),
			},
			{
				Name:        "architecture",
				Description: "The Architecture reported by the node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.NodeInfo.Architecture"),
			},
			{
				Name:          "volumes_in_use",
				Description:   "List of attachable volumes in use (mounted) by the node.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Status.VolumesInUse"),
				IgnoreInTests: true,
			},
			{
				Name:          "config",
				Description:   "Status of the config assigned to the node via the dynamic Kubelet config feature.",
				Type:          schema.TypeJSON,
				Resolver:      resolveCoreNodeConfig,
				IgnoreInTests: true,
			},
			{
				Name:        "hostname",
				Description: "Hostname of the node.",
				Type:        schema.TypeString,
				Resolver:    resolveCoreNodeHostname,
			},
			{
				Name:        "internal_ip",
				Description: "Internal IP address of the node.",
				Type:        schema.TypeInet,
				Resolver:    resolveCoreNodeIP(corev1.NodeInternalIP),
			},
			{
				Name:          "external_ip",
				Description:   "External IP address of the node.",
				Type:          schema.TypeInet,
				Resolver:      resolveCoreNodeIP(corev1.NodeExternalIP),
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_core_node_images",
				Description: "List of container images on this node.",
				Resolver:    fetchCoreNodeImages,
				Columns: []schema.Column{
					{
						Name:        "node_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_nodes table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "names",
						Description: "Names by which this image is known.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "size_bytes",
						Description: "The size of the image in bytes.",
						Type:        schema.TypeBigInt,
					},
				},
			},
			{
				IgnoreInTests: true,
				Name:          "k8s_core_node_volumes_attached",
				Description:   "List of volumes that are attached to the node.",
				Resolver:      fetchCoreNodeVolumesAttached,
				Columns: []schema.Column{
					{
						Name:        "node_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_nodes table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the attached volume.",
						Type:        schema.TypeString,
					},
					{
						Name:        "device_path",
						Description: "Device path where the volume should be available.",
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
func fetchCoreNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	nodes := meta.(*client.Client).Services().Nodes
	opts := metav1.ListOptions{}
	for {
		result, err := nodes.List(ctx, opts)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}

func resolveCoreNodeOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)
	b, err := json.Marshal(node.ObjectMeta.OwnerReferences)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolveCoreNodePodCIDR(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)

	if node.Spec.PodCIDR == "" {
		return nil
	}
	_, n, err := net.ParseCIDR(node.Spec.PodCIDR)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, n))
}

func resolveCoreNodePodCIDRs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)
	cidrs := make([]*net.IPNet, 0, len(node.Spec.PodCIDRs))
	for _, v := range node.Spec.PodCIDRs {
		_, n, err := net.ParseCIDR(v)
		if err != nil {
			return diag.WrapError(err)
		}
		cidrs = append(cidrs, n)
	}
	return diag.WrapError(resource.Set(c.Name, cidrs))
}

func resolveCoreNodeTaints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)
	b, err := json.Marshal(node.Spec.Taints)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolveCoreNodeConditions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)
	b, err := json.Marshal(node.Status.Conditions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func fetchCoreNodeImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	node := parent.Item.(corev1.Node)
	res <- node.Status.Images
	return nil
}

func fetchCoreNodeVolumesAttached(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	node := parent.Item.(corev1.Node)
	res <- node.Status.VolumesAttached
	return nil
}

func resolveCoreNodeConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)
	b, err := json.Marshal(node.Status.Config)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func fetchAddressValue(addrs []corev1.NodeAddress, key corev1.NodeAddressType) (string, bool) {
	for _, a := range addrs {
		if a.Type == key {
			return a.Address, true
		}
	}
	return "", false
}

func resolveCoreNodeHostname(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node := resource.Item.(corev1.Node)
	v, ok := fetchAddressValue(node.Status.Addresses, corev1.NodeHostName)
	if !ok {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, v))
}

func resolveCoreNodeIP(key corev1.NodeAddressType) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		node := resource.Item.(corev1.Node)
		v, ok := fetchAddressValue(node.Status.Addresses, key)
		if !ok {
			return nil
		}
		ip := net.ParseIP(v)
		if ip == nil {
			return diag.WrapError(fmt.Errorf("failed to convert %v to IP address", v))
		}
		if v4 := ip.To4(); v4 != nil {
			ip = v4
		}
		return diag.WrapError(resource.Set(c.Name, ip))
	}
}
