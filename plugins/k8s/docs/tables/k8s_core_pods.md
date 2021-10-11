
# Table: k8s_core_pods
Pod is a collection of containers that can run on a host
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
|restart_policy|text|Restart policy for all containers within the pod.|
|termination_grace_period_seconds|bigint|Optional duration in seconds the pod needs to terminate gracefully|
|active_deadline_seconds|bigint|Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers.|
|dns_policy|text|Sets DNS policy for the pod.|
|node_selector|jsonb|Selector which must be true for the pod to fit on a node.|
|service_account_name|text|Name of the ServiceAccount to use to run this pod.|
|automount_service_account_token|boolean|Indicates whether a service account token should be automatically mounted.|
|node_name|text|Requests to schedule this pod onto a specific node.|
|host_network|boolean|Host networking requested for this pod.|
|host_pid|boolean|Use the host's pid namespace.|
|host_ipc|boolean|Use the host's ipc namespace.|
|share_process_namespace|boolean|Share a single process namespace between all of the containers in a pod.|
|security_context|jsonb|Holds pod-level security attributes and common container settings.|
|image_pull_secrets|jsonb|Optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.|
|hostname|text|Specifies the hostname of the Pod.|
|subdomain|text|Specifies the subdomain of the Pod.|
|affinity|jsonb|If specified, the pod's scheduling constraints.|
|scheduler_name|text|If specified, the pod will be dispatched by specified scheduler.|
|tolerations|jsonb|If specified, the pod's tolerations.|
|host_aliases|jsonb|Optional list of hosts and IPs that will be injected into the pod's hosts file if specified.|
|priority_class_name|text|If specified, indicates the pod's priority|
|priority|integer|The priority value|
|dns_config|jsonb|Specifies the DNS parameters of a pod.|
|readiness_gates|jsonb|If specified, all readiness gates will be evaluated for pod readiness.|
|runtime_class_name|text|Refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod|
|enable_service_links|boolean|Indicates whether information about services should be injected into pod's environment variables.|
|preemption_policy|text|Policy for preempting pods with lower priority.|
|overhead|jsonb|Represents the resource overhead associated with running a pod for a given RuntimeClass.|
|topology_spread_constraints|jsonb|Describes how a group of pods ought to spread across topology domains|
|set_hostname_as_fqdn|boolean|If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name.|
|phase|text|The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle.|
|conditions|jsonb|Current service state of pod.|
|message|text|A human readable message indicating details about why the pod is in this condition.|
|reason|text|A brief CamelCase message indicating details about why the pod is in this state.|
|nominated_node_name|text|Set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods.|
|host_ip|inet|IP address of the host to which the pod is assigned.|
|pod_ip|inet|IP address allocated to the pod.|
|pod_ips|inet[]|podIPs holds the IP addresses allocated to the pod|
|qos_class|text|The Quality of Service (QOS) classification assigned to the pod based on resource requirements.|
