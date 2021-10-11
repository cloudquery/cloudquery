
# Table: k8s_core_nodes
Node is a worker node in Kubernetes.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|k8s_config_context|text|Name of the context from k8s configuration.|
|kind|text|Kind is a string value representing the REST resource this object represents.|
|api_version|text|Defines the versioned schema of this representation of an object.|
|name|text|Unique name within a namespace.|
|namespace|text|Namespace defines the space within which each name must be unique.|
|uid|text|UID is the unique in time and space value for this object.|
|resource_version|text|An opaque value that represents the internal version of this object.|
|generation|bigint|A sequence number representing a specific generation of the desired state.|
|deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate.|
|labels|jsonb|Map of string keys and values that can be used to organize and categorize objects.|
|annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools.|
|owner_references|jsonb|List of objects depended by this object.|
|finalizers|text[]|List of finalizers|
|cluster_name|text|The name of the cluster which the object belongs to.|
|pod_cidr|cidr|Represents the pod IP range assigned to the node.|
|pod_cidrs|cidr[]|Represents the IP ranges assigned to the node for usage by Pods on that node|
|provider_id|text|ID of the node assigned by the cloud provider.|
|unschedulable|boolean|Unschedulable controls node schedulability of new pods|
|taints|jsonb|If specified, the node's taints.|
|capacity|jsonb|Capacity represents the total resources of a node.|
|allocatable|jsonb|Allocatable represents the resources of a node that are available for scheduling.|
|phase|text|NodePhase is the recently observed lifecycle phase of the node.|
|conditions|jsonb|Conditions is an array of current observed node conditions.|
|daemon_endpoints_kubelet_endpoint_port|integer|Port number of the given endpoint.|
|machine_id|text|MachineID reported by the node|
|system_uuid|text|SystemUUID reported by the node|
|boot_id|text|Boot ID reported by the node.|
|kernel_version|text|Kernel Version reported by the node from 'uname -r'|
|os_image|text|OS Image reported by the node from /etc/os-release|
|container_runtime_version|text|Container runtime version reported by the node through runtime remote API.|
|kubelet_version|text|Kubelet Version reported by the node.|
|kube_proxy_version|text|KubeProxy Version reported by the node.|
|operating_system|text|The Operating System reported by the node.|
|architecture|text|The Architecture reported by the node.|
|volumes_in_use|text[]|List of attachable volumes in use (mounted) by the node.|
|config|jsonb|Status of the config assigned to the node via the dynamic Kubelet config feature.|
|hostname|text|Hostname of the node.|
|internal_ip|inet|Internal IP address of the node.|
|external_ip|inet|External IP address of the node.|
