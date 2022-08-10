
# Table: k8s_apps_stateful_sets
StatefulSet represents a set of pods with consistent identities. Identities are defined as:  - Network: A single stable DNS and hostname.  - Storage: As many VolumeClaims as requested. The StatefulSet guarantees that a given network identity will always map to the same storage identity.
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
|replicas|integer|replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. TODO: Consider a rename of this field.|
|selector_match_labels|jsonb|matchLabels is a map of {key,value} pairs|
|template|jsonb|template is the object that describes the pod that will be created if insufficient replicas are detected|
|volume_claim_templates|jsonb|volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod|
|service_name|text|serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set|
|pod_management_policy|text|podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down|
|update_strategy_type|text|Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.|
|update_strategy_rolling_update_partition|integer|Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.|
|revision_history_limit|integer|revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history|
|min_ready_seconds|integer|Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.|
|status_observed_generation|bigint|observedGeneration is the most recent generation observed for this StatefulSet|
|status_replicas|integer|replicas is the number of Pods created by the StatefulSet controller.|
|status_ready_replicas|integer|readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.|
|status_current_replicas|integer|currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.|
|status_updated_replicas|integer|updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.|
|status_current_revision|text|currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).|
|status_update_revision|text|updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)|
|status_collision_count|integer|collisionCount is the count of hash collisions for the StatefulSet|
|status_available_replicas|integer|Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset. This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate. Remove omitempty when graduating to beta|
