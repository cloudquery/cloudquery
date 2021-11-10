
# Table: k8s_batch_cron_jobs
CronJob represents the configuration of a single cron job.
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
|schedule|text|The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.|
|starting_deadline_seconds|bigint|Optional deadline in seconds for starting the job if it misses scheduled time for any reason|
|concurrency_policy|text|Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one|
|suspend|boolean|This flag tells the controller to suspend subsequent executions, it does not apply to already started executions|
|job_template|jsonb|Specifies the job that will be created when executing a CronJob.|
|successful_jobs_history_limit|integer|The number of successful finished jobs to retain|
|failed_jobs_history_limit|integer|The number of failed finished jobs to retain|
|status|jsonb|Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status|
