
# Table: gcp_cloudrun_service_spec_template_containers
A single application container
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_services table (FK)|
|args|text[]|Arguments to the entrypoint|
|command|text[]||
|image|text|Only supports containers from Google Container Registry or Artifact Registry URL of the Container image|
|image_pull_policy|text|Image pull policy|
|liveness_probe_exec_command|text[]|Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem|
|liveness_probe_failure_threshold|bigint|Minimum consecutive failures for the probe to be considered failed after having succeeded|
|liveness_probe_http_get_host|text|Host name to connect to, defaults to the pod IP|
|liveness_probe_http_get_path|text|Path to access on the HTTP server|
|liveness_probe_http_get_scheme|text|Scheme to use for connecting to the host|
|liveness_probe_initial_delay_seconds|bigint|Number of seconds after the container has started before liveness probes are initiated|
|liveness_probe_period_seconds|bigint|How often (in seconds) to perform the probe|
|liveness_probe_success_threshold|bigint|Minimum consecutive successes for the probe to be considered successful after having failed|
|liveness_probe_tcp_socket_host|text|Host name to connect to, defaults to the pod IP|
|liveness_probe_tcp_socket_port|bigint|Number or name of the port to access on the container|
|liveness_probe_timeout_seconds|bigint|Number of seconds after which the probe times out|
|name|text|Name of the container specified as a DNS_LABEL Currently unused in Cloud Run|
|ports|jsonb|List of ports to expose from the container|
|readiness_probe_exec_command|text[]|Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem|
|readiness_probe_failure_threshold|bigint|Minimum consecutive failures for the probe to be considered failed after having succeeded|
|readiness_probe_http_get_host|text|Host name to connect to, defaults to the pod IP|
|readiness_probe_http_get_http_headers|jsonb|Custom headers to set in the request|
|readiness_probe_http_get_path|text|Path to access on the HTTP server|
|readiness_probe_http_get_scheme|text|Scheme to use for connecting to the host|
|readiness_probe_initial_delay_seconds|bigint|Number of seconds after the container has started before liveness probes are initiated|
|readiness_probe_period_seconds|bigint|How often (in seconds) to perform the probe|
|readiness_probe_success_threshold|bigint|Minimum consecutive successes for the probe to be considered successful after having failed|
|readiness_probe_tcp_socket_host|text|Host name to connect to, defaults to the pod IP|
|readiness_probe_tcp_socket_port|bigint|Number or name of the port to access on the container|
|readiness_probe_timeout_seconds|bigint|Number of seconds after which the probe times out|
|resources_limits|jsonb|Only memory and CPU are supported|
|resources_requests|jsonb|Only memory and CPU are supported|
|security_context_run_as_user|bigint|The UID to run the entrypoint of the container process|
|startup_probe_exec_command|text[]|Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem|
|startup_probe_failure_threshold|bigint|Minimum consecutive failures for the probe to be considered failed after having succeeded|
|startup_probe_http_get_host|text|Host name to connect to, defaults to the pod IP|
|startup_probe_http_get_path|text|Path to access on the HTTP server|
|startup_probe_http_get_scheme|text|Scheme to use for connecting to the host|
|startup_probe_initial_delay_seconds|bigint|Number of seconds after the container has started before liveness probes are initiated|
|startup_probe_period_seconds|bigint|How often (in seconds) to perform the probe|
|startup_probe_success_threshold|bigint|Minimum consecutive successes for the probe to be considered successful after having failed|
|startup_probe_tcp_socket_host|text|Host name to connect to, defaults to the pod IP|
|startup_probe_tcp_socket_port|bigint|Number or name of the port to access on the container|
|startup_probe_timeout_seconds|bigint|Number of seconds after which the probe times out|
|termination_message_path|text|Path at which the file to which the container's termination message will be written is mounted into the container's filesystem|
|termination_message_policy|text|Indicate how the termination message should be populated|
|working_dir|text|Container's working directory|
