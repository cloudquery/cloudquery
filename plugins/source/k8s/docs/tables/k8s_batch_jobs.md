
# Table: k8s_batch_jobs
Job represents the configuration of a single job.
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
|parallelism|integer|Specifies the maximum desired number of pods the job should run at any given time|
|completions|integer|Specifies the desired number of successfully finished pods the job should be run with|
|active_deadline_seconds|bigint|Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer|
|backoff_limit|integer|Specifies the number of retries before marking this job failed. Defaults to 6|
|selector_match_labels|jsonb|matchLabels is a map of {key,value} pairs|
|manual_selector|boolean|manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template|
|template|jsonb|Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/|
|ttl_seconds_after_finished|integer|ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed)|
|completion_mode|text|CompletionMode specifies how Pod completions are tracked|
|suspend|boolean|Suspend specifies whether the Job controller should create Pods or not|
|status_active|integer|The number of actively running pods.|
|status_succeeded|integer|The number of pods which reached phase Succeeded.|
|status_failed|integer|The number of pods which reached phase Failed.|
|status_completed_indexes|text|CompletedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format|
|status_uncounted_terminated_pods_succeeded|text[]|Succeeded holds UIDs of succeeded Pods. +listType=set|
|status_uncounted_terminated_pods_failed|text[]|Failed holds UIDs of failed Pods. +listType=set|
