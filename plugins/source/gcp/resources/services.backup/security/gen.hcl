service = "gcp"
output_directory = "."
add_generate     = true

resource "gcp" "security" "secrets" {
  path = "google.golang.org/api/secretmanager/v1.Secret"

  options {
    primary_keys = [
      "resource_name"
    ]
  }
  
  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/plugins/source/gcp/client.ProjectMultiplex"
  }
  
  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/plugins/source/gcp/client.DeleteProjectFilter"
  }

  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreErrorHandler"
  }

  userDefinedColumn "project_id" {
    type = "string"
    description = "GCP Project Id of the resource"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ResolveProject"
    }
  }

  userDefinedColumn "id" {
    type = "string"
    description = "The id of the secret"
    generate_resolver = true
  }

  userDefinedColumn "topics" {
    type = "stringArray"
    description = "A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions. In the format `projects/*/topics/*`"
    generate_resolver = true
  }

  column "name" {
    rename = "resource_name"
    type = "string"
    description = "The resource name of the Secret in the format `projects/*/secrets/*`"
  }

  column "create_time" {
    description = "The time at which the Secret was created"
  }

  column "etag" {
    description = "Etag of the currently stored Secret"
  }

  column "labels" {
    description = "The labels assigned to this Secret"
  }

  column "rotation_next_rotation_time" {
    rename = "next_rotation_time"
    description = "Timestamp in UTC at which the Secret is scheduled to rotate" 
  }

  column "rotation_period" {
    skip = true
  }

  userDefinedColumn "is_automatically_replicated" {
    description = "If true, the secret is automatically replicated by GCP. Otherwise, replications are user-managed."
    type = "bool"
    generate_resolver = true
  }

  column "replication_automatic_customer_managed_encryption_kms_key_name" {
    rename = "automatic_replication_customer_managed_encryption_kms_key_name"
    description = "If the secret is automatically replicated, contains the customer-managed-encryption kms-key-name. Only valid if 'is_automatically_replicated' is true. If null, then the secret is encrypted with a google-managed key."
    generate_resolver = true
  }

  column "ttl" {
    skip = true
  }

  column "expire_time" {
    description = "Timestamp in UTC when the Secret is scheduled to expire."
  }

  relation "gcp" "security" "replication_user_managed_replicas" {
    rename = "user_managed_replicas"
    description = "Describes user-managed replicas of this secret. Empty for automatically replicated secrets"

    column "customer_managed_encryption_kms_key_name" {
      description = "If the replica is encrypted with customer-managed encryption, contains the kms key name. If the column is NULL, the replica is encrypted with a google-managed key"
    }

    column "location" {
      description = "The canonical IDs of the location to replicate data. For example: \"us-east1\""
    }
  }

}
