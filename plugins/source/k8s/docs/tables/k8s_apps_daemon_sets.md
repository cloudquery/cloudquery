
# Table: k8s_apps_daemon_sets
DaemonSet represents the configuration of a daemon set.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|context|text|Name of the context from k8s configuration.|
|name|text|Name must be unique within a namespace|
|generate_name|text|GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed|
|namespace|text|Namespace defines the space within which each name must be unique|
|self_link|text|SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.|
|uid|text|UID is the unique in time and space value for this object|
|resource_version|text|An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed|
|generation|bigint|A sequence number representing a specific generation of the desired state. Populated by the system|
|deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate before it will be removed from the system|
|labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|owner_references|jsonb|List of objects depended by this object|
|finalizers|text[]|Must be empty before the object is deleted from the registry|
|cluster_name|text|The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.|
|managed_fields|jsonb|ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow|
|selector_match_labels|jsonb|matchLabels is a map of {key,value} pairs|
|template|jsonb|An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template|
|update_strategy_type|text|Type of daemon set update|
|update_strategy_rolling_update_max_unavailable_type|bigint||
|update_strategy_rolling_update_max_unavailable_int_val|integer||
|update_strategy_rolling_update_max_unavailable_str_val|text||
|update_strategy_rolling_update_max_surge_type|bigint||
|update_strategy_rolling_update_max_surge_int_val|integer||
|update_strategy_rolling_update_max_surge_str_val|text||
|min_ready_seconds|integer|The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available|
|revision_history_limit|integer|The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.|
|status_current_number_scheduled|integer|The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/|
|status_number_misscheduled|integer|The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/|
|status_desired_number_scheduled|integer|The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/|
|status_number_ready|integer|The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.|
|status_observed_generation|bigint|The most recent generation observed by the daemon set controller.|
|status_updated_number_scheduled|integer|The total number of nodes that are running updated daemon pod|
|status_number_available|integer|The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)|
|status_number_unavailable|integer|The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)|
|status_collision_count|integer|Count of hash collisions for the DaemonSet|
