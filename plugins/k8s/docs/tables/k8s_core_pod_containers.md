
# Table: k8s_core_pod_containers
A single application container that you want to run within a pod.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pod_cq_id|uuid|Unique CloudQuery ID of k8s_core_pods table (FK)|
|name|text|Name of the container specified as a DNS_LABEL.|
|image|text|Docker image name.|
|command|text[]|Entrypoint array|
|args|text[]|Arguments to the entrypoint.|
|working_dir|text|Container's working directory.|
|env_from|jsonb|List of sources to populate environment variables in the container.|
|resources_limits|jsonb|Limits describes the maximum amount of compute resources allowed.|
|resources_requests|jsonb|Requests describes the minimum amount of compute resources required.|
|liveness_probe|jsonb|Periodic probe of container liveness.|
|readiness_probe|jsonb|Periodic probe of container service readiness.|
|startup_probe|jsonb|Startup probe indicates that the Pod has successfully initialized.|
|lifecycle|jsonb|Actions that the management system should take in response to container lifecycle events.|
|termination_message_path|text|Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.|
|termination_message_policy|text|Indicate how the termination message should be populated.|
|image_pull_policy|text|Image pull policy.|
|security_context|jsonb|security options the container should be run with.|
|stdin|boolean|Whether this container should allocate a buffer for stdin in the container runtime|
|stdin_once|boolean|Whether the container runtime should close the stdin channel after it has been opened by a single attach|
|tty|boolean|Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false. +optional|
