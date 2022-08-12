package cloudrun

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	run "google.golang.org/api/run/v1"
)

//go:generate cq-gen --resource services --config gen.hcl --output .
func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudrun_services",
		Description: "Service acts as a top-level container that manages a set of Routes and Configurations which implement a network service",
		Resolver:    fetchCloudrunServices,
		Multiplex:   client.ProjectMultiplex,

		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:        "api_version",
				Description: "The API version for this call such as \"servingknativedev/v1\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "The kind of resource, in this case \"Service\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "metadata_annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Metadata.Annotations"),
			},
			{
				Name:        "metadata_creation_timestamp",
				Description: "CreationTimestamp is a timestamp representing the server time when this object was created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.CreationTimestamp"),
			},
			{
				Name:        "metadata_generation",
				Description: "A sequence number representing a specific generation of the desired state",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Metadata.Generation"),
			},
			{
				Name:        "metadata_labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Metadata.Labels"),
			},
			{
				Name:        "metadata_name",
				Description: "Name must be unique within a namespace, within a Cloud Run region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Name"),
			},
			{
				Name:        "metadata_namespace",
				Description: "Namespace defines the space within each name must be unique, within a Cloud Run region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Namespace"),
			},
			{
				Name:        "metadata_resource_version",
				Description: "Optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.ResourceVersion"),
			},
			{
				Name:        "metadata_self_link",
				Description: "SelfLink is a URL representing this object Populated by the system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.SelfLink"),
			},
			{
				Name:        "metadata_uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Uid"),
			},
			{
				Name:        "spec_template_metadata_annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.Annotations"),
			},
			{
				Name:        "spec_template_metadata_creation_timestamp",
				Description: "CreationTimestamp is a timestamp representing the server time when this object was created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.CreationTimestamp"),
			},
			{
				Name:        "spec_template_metadata_generation",
				Description: "A sequence number representing a specific generation of the desired state",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.Generation"),
			},
			{
				Name:        "spec_template_metadata_labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.Labels"),
			},
			{
				Name:        "spec_template_metadata_name",
				Description: "Name must be unique within a namespace, within a Cloud Run region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.Name"),
			},
			{
				Name:        "spec_template_metadata_namespace",
				Description: "Namespace defines the space within each name must be unique, within a Cloud Run region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.Namespace"),
			},
			{
				Name:        "spec_template_metadata_resource_version",
				Description: "Optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.ResourceVersion"),
			},
			{
				Name:        "spec_template_metadata_self_link",
				Description: "SelfLink is a URL representing this object Populated by the system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.SelfLink"),
			},
			{
				Name:        "spec_template_metadata_uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Metadata.Uid"),
			},
			{
				Name:        "spec_template_spec_container_concurrency",
				Description: "Optional",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.Spec.ContainerConcurrency"),
			},
			{
				Name:        "spec_template_spec_service_account_name",
				Description: "Email address of the IAM service account associated with the revision of the service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Template.Spec.ServiceAccountName"),
			},
			{
				Name:        "spec_template_spec_timeout_seconds",
				Description: "TimeoutSeconds holds the max duration the instance is allowed for responding to a request",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.Template.Spec.TimeoutSeconds"),
			},
			{
				Name:     "status_address_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Address.Url"),
			},
			{
				Name:        "status_latest_created_revision_name",
				Description: "From ConfigurationStatus LatestCreatedRevisionName is the last revision that was created from this Service's Configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.LatestCreatedRevisionName"),
			},
			{
				Name:        "status_latest_ready_revision_name",
				Description: "From ConfigurationStatus LatestReadyRevisionName holds the name of the latest Revision stamped out from this Service's Configuration that has had its \"Ready\" condition become \"True\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.LatestReadyRevisionName"),
			},
			{
				Name:        "status_observed_generation",
				Description: "ObservedGeneration is the 'Generation' of the Route that was last processed by the controller",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_url",
				Description: "From RouteStatus",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Url"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "gcp_cloudrun_service_metadata_owner_references",
				Description:   "OwnerReference contains enough information to let you identify an owning object",
				Resolver:      fetchCloudrunServiceMetadataOwnerReferences,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_version",
						Description: "API version of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "block_owner_deletion",
						Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed",
						Type:        schema.TypeBool,
					},
					{
						Name:        "controller",
						Description: "If true, this reference points to the managing controller",
						Type:        schema.TypeBool,
					},
					{
						Name:        "kind",
						Description: "Kind of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "uid",
						Description: "UID of the referent",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_cloudrun_service_spec_template_metadata_owner_references",
				Description:   "OwnerReference contains enough information to let you identify an owning object",
				Resolver:      fetchCloudrunServiceSpecTemplateMetadataOwnerReferences,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_version",
						Description: "API version of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "block_owner_deletion",
						Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed",
						Type:        schema.TypeBool,
					},
					{
						Name:        "controller",
						Description: "If true, this reference points to the managing controller",
						Type:        schema.TypeBool,
					},
					{
						Name:        "kind",
						Description: "Kind of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "uid",
						Description: "UID of the referent",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_spec_template_containers",
				Description: "A single application container",
				Resolver:    fetchCloudrunServiceSpecTemplateContainers,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "args",
						Description:   "Arguments to the entrypoint",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
					{
						Name: "command",
						Type: schema.TypeStringArray,
					},
					{
						Name:        "image",
						Description: "Only supports containers from Google Container Registry or Artifact Registry URL of the Container image",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_pull_policy",
						Description: "Image pull policy",
						Type:        schema.TypeString,
					},
					{
						Name:          "liveness_probe_exec_command",
						Description:   "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("LivenessProbe.Exec.Command"),
						IgnoreInTests: true,
					},
					{
						Name:        "liveness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.FailureThreshold"),
					},
					{
						Name:        "liveness_probe_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.HttpGet.Host"),
					},
					{
						Name:        "liveness_probe_http_get_path",
						Description: "Path to access on the HTTP server",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.HttpGet.Path"),
					},
					{
						Name:        "liveness_probe_http_get_scheme",
						Description: "Scheme to use for connecting to the host",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.HttpGet.Scheme"),
					},
					{
						Name:        "liveness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "liveness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.PeriodSeconds"),
					},
					{
						Name:        "liveness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.SuccessThreshold"),
					},
					{
						Name:        "liveness_probe_tcp_socket_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LivenessProbe.TcpSocket.Host"),
					},
					{
						Name:        "liveness_probe_tcp_socket_port",
						Description: "Number or name of the port to access on the container",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.TcpSocket.Port"),
					},
					{
						Name:        "liveness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LivenessProbe.TimeoutSeconds"),
					},
					{
						Name:        "name",
						Description: "Name of the container specified as a DNS_LABEL Currently unused in Cloud Run",
						Type:        schema.TypeString,
					},
					{
						Name:        "ports",
						Description: "List of ports to expose from the container",
						Type:        schema.TypeJSON,
					},
					{
						Name:          "readiness_probe_exec_command",
						Description:   "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("ReadinessProbe.Exec.Command"),
						IgnoreInTests: true,
					},
					{
						Name:        "readiness_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.FailureThreshold"),
					},
					{
						Name:        "readiness_probe_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.HttpGet.Host"),
					},
					{
						Name:          "readiness_probe_http_get_http_headers",
						Description:   "Custom headers to set in the request",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("ReadinessProbe.HttpGet.HttpHeaders"),
						IgnoreInTests: true,
					},
					{
						Name:        "readiness_probe_http_get_path",
						Description: "Path to access on the HTTP server",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.HttpGet.Path"),
					},
					{
						Name:        "readiness_probe_http_get_scheme",
						Description: "Scheme to use for connecting to the host",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.HttpGet.Scheme"),
					},
					{
						Name:        "readiness_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.InitialDelaySeconds"),
					},
					{
						Name:        "readiness_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.PeriodSeconds"),
					},
					{
						Name:        "readiness_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.SuccessThreshold"),
					},
					{
						Name:        "readiness_probe_tcp_socket_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReadinessProbe.TcpSocket.Host"),
					},
					{
						Name:        "readiness_probe_tcp_socket_port",
						Description: "Number or name of the port to access on the container",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.TcpSocket.Port"),
					},
					{
						Name:        "readiness_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReadinessProbe.TimeoutSeconds"),
					},
					{
						Name:        "resources_limits",
						Description: "Only memory and CPU are supported",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Resources.Limits"),
					},
					{
						Name:          "resources_requests",
						Description:   "Only memory and CPU are supported",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Resources.Requests"),
						IgnoreInTests: true,
					},
					{
						Name:        "security_context_run_as_user",
						Description: "The UID to run the entrypoint of the container process",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SecurityContext.RunAsUser"),
					},
					{
						Name:          "startup_probe_exec_command",
						Description:   "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("StartupProbe.Exec.Command"),
						IgnoreInTests: true,
					},
					{
						Name:        "startup_probe_failure_threshold",
						Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.FailureThreshold"),
					},
					{
						Name:        "startup_probe_http_get_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.HttpGet.Host"),
					},
					{
						Name:        "startup_probe_http_get_path",
						Description: "Path to access on the HTTP server",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.HttpGet.Path"),
					},
					{
						Name:        "startup_probe_http_get_scheme",
						Description: "Scheme to use for connecting to the host",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.HttpGet.Scheme"),
					},
					{
						Name:        "startup_probe_initial_delay_seconds",
						Description: "Number of seconds after the container has started before liveness probes are initiated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.InitialDelaySeconds"),
					},
					{
						Name:        "startup_probe_period_seconds",
						Description: "How often (in seconds) to perform the probe",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.PeriodSeconds"),
					},
					{
						Name:        "startup_probe_success_threshold",
						Description: "Minimum consecutive successes for the probe to be considered successful after having failed",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.SuccessThreshold"),
					},
					{
						Name:        "startup_probe_tcp_socket_host",
						Description: "Host name to connect to, defaults to the pod IP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StartupProbe.TcpSocket.Host"),
					},
					{
						Name:        "startup_probe_tcp_socket_port",
						Description: "Number or name of the port to access on the container",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.TcpSocket.Port"),
					},
					{
						Name:        "startup_probe_timeout_seconds",
						Description: "Number of seconds after which the probe times out",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StartupProbe.TimeoutSeconds"),
					},
					{
						Name:        "termination_message_path",
						Description: "Path at which the file to which the container's termination message will be written is mounted into the container's filesystem",
						Type:        schema.TypeString,
					},
					{
						Name:        "termination_message_policy",
						Description: "Indicate how the termination message should be populated",
						Type:        schema.TypeString,
					},
					{
						Name:        "working_dir",
						Description: "Container's working directory",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_cloudrun_service_spec_template_container_env",
						Description: "EnvVar represents an environment variable present in a Container",
						Resolver:    fetchCloudrunServiceSpecTemplateContainerEnvs,
						Columns: []schema.Column{
							{
								Name:        "service_spec_template_container_cq_id",
								Description: "Unique CloudQuery ID of gcp_cloudrun_service_spec_template_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Name of the environment variable",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables",
								Type:        schema.TypeString,
							},
							{
								Name:        "value_from_config_map_key_ref_key",
								Description: "The key to select",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
							},
							{
								Name:        "value_from_config_map_key_ref_local_object_reference_name",
								Description: "Name of the referent",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_config_map_key_ref_name",
								Description: "The ConfigMap to select from",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Name"),
							},
							{
								Name:        "value_from_config_map_key_ref_optional",
								Description: "Specify whether the ConfigMap or its key must be defined",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
							},
							{
								Name:        "value_from_secret_key_ref_key",
								Description: "A Cloud Secret Manager secret version",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
							},
							{
								Name:        "value_from_secret_key_ref_local_object_reference_name",
								Description: "Name of the referent",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
							},
							{
								Name:        "value_from_secret_key_ref_name",
								Description: "The name of the secret in Cloud Secret Manager",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Name"),
							},
							{
								Name:        "value_from_secret_key_ref_optional",
								Description: "Specify whether the Secret or its key must be defined",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
							},
						},
					},
					{
						Name:        "gcp_cloudrun_service_spec_template_container_volume_mounts",
						Description: "VolumeMount describes a mounting of a Volume within a container",
						Resolver:    fetchCloudrunServiceSpecTemplateContainerVolumeMounts,
						Columns: []schema.Column{
							{
								Name:        "service_spec_template_container_cq_id",
								Description: "Unique CloudQuery ID of gcp_cloudrun_service_spec_template_containers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "mount_path",
								Description: "Path within the container at which the volume should be mounted",
								Type:        schema.TypeString,
							},
							{
								Name:        "name",
								Description: "The name of the volume",
								Type:        schema.TypeString,
							},
							{
								Name:        "read_only",
								Description: "Only true is accepted",
								Type:        schema.TypeBool,
							},
							{
								Name:        "sub_path",
								Description: "Path within the volume from which the container's volume should be mounted",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_spec_template_volumes",
				Description: "Volume represents a named volume in a container",
				Resolver:    fetchCloudrunServiceSpecTemplateVolumes,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "config_map_default_mode",
						Description: "Integer representation of mode bits to use on created files by default",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ConfigMap.DefaultMode"),
					},
					{
						Name:        "config_map_name",
						Description: "Name of the config",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigMap.Name"),
					},
					{
						Name:        "config_map_optional",
						Description: "Specify whether the Secret or its keys must be defined",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ConfigMap.Optional"),
					},
					{
						Name:        "name",
						Description: "Volume's name",
						Type:        schema.TypeString,
					},
					{
						Name:        "secret_default_mode",
						Description: "Integer representation of mode bits to use on created files by default",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Secret.DefaultMode"),
					},
					{
						Name:        "secret_optional",
						Description: "Specify whether the Secret or its keys must be defined",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Secret.Optional"),
					},
					{
						Name:        "secret_name",
						Description: "The name of the secret in Cloud Secret Manager",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Secret.SecretName"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "gcp_cloudrun_service_spec_template_volume_config_map_items",
						Description:   "Maps a string key to a path within a volume",
						Resolver:      fetchCloudrunServiceSpecTemplateVolumeConfigMapItems,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "service_spec_template_volume_cq_id",
								Description: "Unique CloudQuery ID of gcp_cloudrun_service_spec_template_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key",
								Description: "The Cloud Secret Manager secret version",
								Type:        schema.TypeString,
							},
							{
								Name:        "mode",
								Description: "Mode bits to use on this file, must be a value between 01 and 0777 (octal)",
								Type:        schema.TypeBigInt,
							},
							{
								Name:        "path",
								Description: "The relative path of the file to map the key to",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "gcp_cloudrun_service_spec_template_volume_secret_items",
						Description: "Maps a string key to a path within a volume",
						Resolver:    fetchCloudrunServiceSpecTemplateVolumeSecretItems,
						Columns: []schema.Column{
							{
								Name:        "service_spec_template_volume_cq_id",
								Description: "Unique CloudQuery ID of gcp_cloudrun_service_spec_template_volumes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key",
								Description: "The Cloud Secret Manager secret version",
								Type:        schema.TypeString,
							},
							{
								Name:        "mode",
								Description: "Mode bits to use on this file, must be a value between 01 and 0777 (octal)",
								Type:        schema.TypeBigInt,
							},
							{
								Name:        "path",
								Description: "The relative path of the file to map the key to",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_spec_traffic",
				Description: "TrafficTarget holds a single entry of the routing table for a Route",
				Resolver:    fetchCloudrunServiceSpecTraffics,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "configuration_name",
						Description: "ConfigurationName of a configuration to whose latest revision we will send this portion of traffic",
						Type:        schema.TypeString,
					},
					{
						Name:        "latest_revision",
						Description: "Optional",
						Type:        schema.TypeBool,
					},
					{
						Name:        "percent",
						Description: "Percent specifies percent of the traffic to this Revision or Configuration",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "revision_name",
						Description: "RevisionName of a specific revision to which to send this portion of traffic",
						Type:        schema.TypeString,
					},
					{
						Name:        "tag",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "url",
						Description: "Output only",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_status_conditions",
				Description: "Condition defines a generic condition for a Resource",
				Resolver:    fetchCloudrunServiceStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "last_transition_time",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "severity",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "type is used to communicate the status of the reconciliation process",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_status_traffic",
				Description: "TrafficTarget holds a single entry of the routing table for a Route",
				Resolver:    fetchCloudrunServiceStatusTraffics,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "configuration_name",
						Description: "ConfigurationName of a configuration to whose latest revision we will send this portion of traffic",
						Type:        schema.TypeString,
					},
					{
						Name:        "latest_revision",
						Description: "Optional",
						Type:        schema.TypeBool,
					},
					{
						Name:        "percent",
						Description: "Percent specifies percent of the traffic to this Revision or Configuration",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "revision_name",
						Description: "RevisionName of a specific revision to which to send this portion of traffic",
						Type:        schema.TypeString,
					},
					{
						Name:        "tag",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "url",
						Description: "Output only",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCloudrunServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	var nextPageToken string
	for {
		if c.Services == nil || c.Services.CloudRun == nil || c.Services.CloudRun.Projects == nil || c.Services.CloudRun.Projects.Locations == nil || c.Services.CloudRun.Projects.Locations.Services == nil {
			return nil
		}
		call := c.Services.CloudRun.Projects.Locations.Services.List("projects/" + c.ProjectId + "/locations/-").Continue(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return errors.WithStack(err)
		}
		if output.Items == nil {
			return nil
		}

		res <- output.Items
		if output.Metadata == nil || output.Metadata.Continue == "" {
			break
		}
		nextPageToken = output.Metadata.Continue
	}
	return nil
}
func fetchCloudrunServiceMetadataOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Metadata == nil {
		return nil
	}
	res <- r.Metadata.OwnerReferences
	return nil
}
func fetchCloudrunServiceSpecTemplateMetadataOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Spec == nil || r.Spec.Template == nil || r.Spec.Template.Metadata == nil {
		return nil
	}
	res <- r.Spec.Template.Metadata.OwnerReferences
	return nil
}
func fetchCloudrunServiceSpecTemplateContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Spec == nil || r.Spec.Template == nil || r.Spec.Template.Spec == nil {
		return nil
	}
	res <- r.Spec.Template.Spec.Containers
	return nil
}
func fetchCloudrunServiceSpecTemplateContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	if r == nil {
		return nil
	}
	res <- r.Env
	return nil
}
func fetchCloudrunServiceSpecTemplateContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	if r == nil {
		return nil
	}
	res <- r.VolumeMounts
	return nil
}
func fetchCloudrunServiceSpecTemplateVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Spec == nil || r.Spec.Template == nil || r.Spec.Template.Spec == nil {
		return nil
	}
	res <- r.Spec.Template.Spec.Volumes
	return nil
}
func fetchCloudrunServiceSpecTemplateVolumeConfigMapItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Volume)
	if r == nil || r.ConfigMap == nil {
		return nil
	}
	res <- r.ConfigMap.Items
	return nil
}
func fetchCloudrunServiceSpecTemplateVolumeSecretItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Volume)
	if r == nil || r.Secret == nil {
		return nil
	}
	res <- r.Secret.Items
	return nil
}
func fetchCloudrunServiceSpecTraffics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Spec == nil {
		return nil
	}
	res <- r.Spec.Traffic
	return nil
}
func fetchCloudrunServiceStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Status == nil {
		return nil
	}
	res <- r.Status.Conditions
	return nil
}
func fetchCloudrunServiceStatusTraffics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r == nil || r.Status == nil {
		return nil
	}
	res <- r.Status.Traffic
	return nil
}
