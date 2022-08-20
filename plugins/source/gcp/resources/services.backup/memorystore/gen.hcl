service          = "gcp"
output_directory = "."
add_generate     = true

resource "gcp" "memorystore" "redis_instances" {
  path        = "google.golang.org/api/redis/v1.Instance"
  description = "A Memorystore for Redis instance"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreErrorHandler"
  }

  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/plugins/source/gcp/client.ProjectMultiplex"
  }

  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/plugins/source/gcp/client.DeleteProjectFilter"
  }

  options {
    primary_keys = [
      "project_id",
      "id"
    ]
  }

  ignore_columns_in_tests = [
    "suspension_reasons",
    "maintenance_policy",
    "maintenance_schedule",
    "redis_configs",
  ]

  userDefinedColumn "project_id" {
    type        = "string"
    description = "GCP Project ID of the resource"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ResolveProject"
    }
  }

  userDefinedColumn "id" {
    type              = "string"
    description       = "Memorystore for Redis instance ID"
    generate_resolver = true
  }

  column "alternative_location_id" {
    description = "If specified, at least one node will be provisioned in this zone in addition to the zone specified in `location_id`"
  }

  column "auth_enabled" {
    description = "Indicates whether OSS Redis AUTH is enabled for the instance"
  }

  column "authorized_network" {
    description = "The full name of the Google Compute Engine network (https://cloud.google.com/vpc/docs/vpc) to which the instance is connected"
  }

  column "connect_mode" {
    description = "The network connect mode of the Redis instance"
  }

  column "create_time" {
    description = "The time the instance was created"
  }

  column "current_location_id" {
    description = "The current zone where the Redis primary node is located"
  }

  column "customer_managed_key" {
    description = "The KMS key reference that the customer provides when trying to create the instance"
  }

  column "display_name" {
    description = "An arbitrary and optional user-provided name for the instance"
  }

  column "host" {
    description = "Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service"
  }

  column "labels" {
    description = "Resource labels to represent user provided metadata"
  }

  column "location_id" {
    description = "The zone where the instance will be provisioned"
  }

  column "maintenance_policy" {
    description = "The maintenance policy for the instance"
    type        = "json"
  }

  column "maintenance_schedule" {
    description = "Date and time of upcoming maintenance events which have been scheduled"
    type        = "json"
  }

  column "memory_size_gb" {
    description = "Redis memory size in GiB"
  }

  column "name" {
    description = "Unique name of the resource"
  }

  column "persistence_config" {
    description = "Persistence configuration parameters"
    type        = "json"
  }

  column "persistence_iam_identity" {
    description = "Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage"
  }

  column "port" {
    description = "The port number of the exposed Redis endpoints"
  }

  column "read_endpoint" {
    description = "Hostname or IP address of the exposed readonly Redis endpoint"
  }

  column "read_endpoint_port" {
    description = "The port number of the exposed readonly redis endpoint"
  }

  column "read_replicas_mode" {
    description = "Read replicas mode for the instance"
  }

  column "redis_configs" {
    description = "Redis configuration parameters"
  }

  column "redis_version" {
    description = "The version of Redis software"
  }

  column "replica_count" {
    description = "The number of replica nodes"
  }

  column "reserved_ip_range" {
    description = "IP range for node placement"
  }

  column "secondary_ip_range" {
    description = "Additional IP range for node placement"
  }

  column "state" {
    description = "The current state of the instance"
  }

  column "status_message" {
    description = "Additional information about the current status of the instance, if available"
  }

  column "suspension_reasons" {
    description = "Reasons that causes instance in `SUSPENDED` state"
  }

  column "tier" {
    description = "The service tier of the instance"
  }

  column "transit_encryption_mode" {
    description = "The TLS mode of the Redis instance"
  }

  column "nodes" {
    description = "Redis instance nodes properties"
    type        = "json"
  }

  relation "gcp" "memorystore_redis_instance" "server_ca_certs" {
    description = "List of server CA certificates for the instance"

    options {
      primary_keys = ["sha1_fingerprint"]
    }

    column "cert" {
      description = "PEM representation"
    }

    column "create_time" {
      description = "The time when the certificate was created in RFC 3339 format"
    }

    column "expire_time" {
      description = "The time when the certificate expires in RFC 3339 format"
    }

    column "serial_number" {
      description = "Serial number, as extracted from the certificate"
    }

    column "sha1_fingerprint" {
      description = "Sha1 Fingerprint of the certificate"
    }
  }
}