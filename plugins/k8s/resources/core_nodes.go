package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CoreNodes() *schema.Table {
	return &schema.Table{
		Name:        "k8s_core_nodes",
		Description: "Node is a worker node in Kubernetes.",
		Resolver:    fetchCoreNodes,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
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
				Name:        "deletion_grace_period_seconds",
				Description: "Number of seconds allowed for this object to gracefully terminate.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
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
				Name:        "owner_references",
				Description: "List of objects depended by this object.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreNodeOwnerReferences,
			},
			{
				Name:        "finalizers",
				Description: "List of finalizers",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
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
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.PodCIDR"),
			},
			{
				Name:        "pod_cidrs",
				Description: "Represents the IP ranges assigned to the node for usage by Pods on that node",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Spec.PodCIDRs"),
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
				Name:        "taints",
				Description: "If specified, the node's taints.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreNodeTaints,
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
				Name:        "images",
				Description: "List of container images on this node.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreNodeImages,
			},
			{
				Name:        "volumes_in_use",
				Description: "List of attachable volumes in use (mounted) by the node.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Status.VolumesInUse"),
			},
			{
				Name:        "volumes_attached",
				Description: "List of volumes that are attached to the node.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreNodeVolumesAttached,
			},
			{
				Name:        "config",
				Description: "Status of the config assigned to the node via the dynamic Kubelet config feature.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCoreNodeConfig,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_core_node_addresses",
				Description: "NodeAddress contains information for the node's address.",
				Resolver:    fetchCoreNodeAddresses,
				Columns: []schema.Column{
					{
						Name:        "node_cq_id",
						Description: "Unique CloudQuery ID of k8s_core_nodes table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Node address type, one of Hostname, ExternalIP or InternalIP.",
						Type:        schema.TypeString,
					},
					{
						Name:        "address",
						Description: "The node address.",
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
func fetchCoreNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client).Services.Nodes
	result, err := client.List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	res <- result.Items
	return nil
}

func resolveCoreNodeOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node, ok := resource.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", resource.Item)
	}
	b, err := json.Marshal(node.ObjectMeta.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreNodeTaints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node, ok := resource.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", resource.Item)
	}
	b, err := json.Marshal(node.Spec.Taints)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreNodeConditions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node, ok := resource.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", resource.Item)
	}
	b, err := json.Marshal(node.Status.Conditions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreNodeImages(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node, ok := resource.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", resource.Item)
	}
	b, err := json.Marshal(node.Status.Images)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveCoreNodeVolumesAttached(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node, ok := resource.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", resource.Item)
	}
	b, err := json.Marshal(node.Status.VolumesAttached)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchCoreNodeAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	node, ok := parent.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", parent.Item)
	}
	res <- node.Status.Addresses
	return nil
}

func resolveCoreNodeConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	node, ok := resource.Item.(corev1.Node)
	if !ok {
		return fmt.Errorf("not a corev1.Node instance: %T", resource.Item)
	}
	b, err := json.Marshal(node.Status.Config)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
