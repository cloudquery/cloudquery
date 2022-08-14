
# Table: k8s_core_services
Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|context|text|Name of the context from k8s configuration.|
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
|selector|jsonb|Route service traffic to pods with label keys and values matching this selector|
|cluster_ip|inet|clusterIP is the IP address of the service and is usually assigned randomly|
|cluster_ips|inet[]|ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly|
|type|text|type determines how the Service is exposed|
|external_ips|inet[]|externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service|
|session_affinity|text|Used to maintain session affinity.|
|load_balancer_ip|text|Load balancer will get created with the IP specified in this field.|
|load_balancer_source_ranges|text[]|If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer to the specified client IPs|
|external_name|text|The external reference that discovery mechanisms will return as an alias for this service.|
|external_traffic_policy|text|Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints|
|health_check_node_port|integer|Specifies the healthcheck nodePort for the service.|
|publish_not_ready_addresses|boolean|Indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready.|
|session_affinity_config_client_ip_timeout_seconds|integer|Specifies the seconds of ClientIP type session sticky time.|
|ip_families|text[]|A list of IP families assigned to this service.|
|ip_family_policy|text|Represents the dual-stack-ness requested or required by this Service.|
|allocate_load_balancer_node_ports|boolean|Defines if NodePorts will be automatically allocated for services with type LoadBalancer|
|load_balancer_class|text|The class of the load balancer implementation this Service belongs to.|
|internal_traffic_policy|text|Specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. "Cluster" routes internal traffic to a Service to all endpoints. "Local" routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready.|
