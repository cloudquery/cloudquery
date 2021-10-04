
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
|restart_policy|text|Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy +optional|
|termination_grace_period_seconds|bigint|Optional duration in seconds the pod needs to terminate gracefully|
|active_deadline_seconds|bigint|Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer. +optional|
|dns_policy|text|Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. +optional|
|node_selector|jsonb|NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ +optional +mapType=atomic|
|service_account_name|text|ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/ +optional|
|automount_service_account_token|boolean|AutomountServiceAccountToken indicates whether a service account token should be automatically mounted. +optional|
|node_name|text|NodeName is a request to schedule this pod onto a specific node|
|host_network|boolean|Host networking requested for this pod|
|host_p_id|boolean|Use the host's pid namespace. Optional: Default to false. +k8s:conversion-gen=false +optional|
|host_ip_c|boolean|Use the host's ipc namespace. Optional: Default to false. +k8s:conversion-gen=false +optional|
|share_process_namespace|boolean|Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. +k8s:conversion-gen=false +optional|
|security_context|jsonb|SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty|
|image_pull_secrets|jsonb|ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use|
|hostname|text|Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value. +optional|
|subdomain|text|If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all. +optional|
|affinity|jsonb|If specified, the pod's scheduling constraints +optional|
|scheduler_name|text|If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler. +optional|
|tolerations|jsonb|If specified, the pod's tolerations. +optional|
|host_aliases|jsonb|HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified|
|priority_class_name|text|If specified, indicates the pod's priority|
|priority|integer|The priority value|
|dns_config|jsonb|Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. +optional|
|readiness_gates|jsonb|If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates +optional|
|runtime_class_name|text|RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod|
|enable_service_links|boolean|EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true. +optional|
|preemption_policy|text|PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate. +optional|
|overhead|jsonb|Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller|
|topology_spread_constraints|jsonb|TopologySpreadConstraints describes how a group of pods ought to spread across topology domains|
|set_hostname_as_fqdn|boolean|If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false. +optional|
|phase|text|The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod's status. There are five possible phase values:  Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created|
|conditions|jsonb|Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions +optional +patchMergeKey=type +patchStrategy=merge|
|message|text|A human readable message indicating details about why the pod is in this condition. +optional|
|reason|text|A brief CamelCase message indicating details about why the pod is in this state. e.g|
|nominated_node_name|text|nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node|
|host_ip|text|IP address of the host to which the pod is assigned|
|pod_ip|text|IP address allocated to the pod|
|pod_ips|text[]|podIPs holds the IP addresses allocated to the pod|
|qos_class|text|The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md +optional|
