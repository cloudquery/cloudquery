service          = "gcp"
output_directory = "."
add_generate     = true



description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}

resource "gcp" "compute" "instance_groups" {
  path = "google.golang.org/api/compute/v1.InstanceGroup"
  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/plugins/source/gcp/client.ProjectMultiplex"
  }

  multiplex "DeleteFilter" {
    path = "github.com/cloudquery/plugins/source/gcp/client.DeleteFilter"
  }

  multiplex "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreError"
  }


  options {
    primary_keys = [
      "project_id",
      "id"
    ]
  }

  column "creation_timestamp" {
    type = "timestamp"
    resolver "dateResolver" {
      path          = "github.com/cloudquery/plugin-sdk/schema.DateResolver"
      path_resolver = true
    }
  }

  userDefinedColumn "project_id" {
    type        = "string"
    description = "GCP Project Id of the resource"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ResolveProject"
    }
  }

  column "named_ports" {
    type              = "json"
    generate_resolver = true
  }

  user_relation "gcp" "compute" "instances" {
    path = "google.golang.org/api/compute/v1.InstanceWithNamedPorts"

    column "named_ports" {
      type              = "json"
      generate_resolver = true
    }

    column "status" {
      description = "Status of the instance. One of DEPROVISIONING, PROVISIONING, REPAIRING, RUNNING, STAGING, STOPPED, STOPPING, SUSPENDED, SUSPENDING, TERMINATED"
    }
  }
}