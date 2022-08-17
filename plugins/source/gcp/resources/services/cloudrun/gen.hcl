service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_optional" {
  words = ["(Optional) "]
}

description_modifier "remove_field_name" {
  regex = "^.+: "
}

resource "gcp" "cloudrun" "services" {
  path = "google.golang.org/api/run/v1.Service"

  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/plugins/source/gcp/client.ProjectMultiplex"
  }
  deleteFilter "DeleteFilter" {
    path = "github.com/cloudquery/plugins/source/gcp/client.DeleteProjectFilter"
  }
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreErrorHandler"
  }

  userDefinedColumn "project_id" {
    type = "string"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ResolveProject"
    }
  }

  column "create_time" {
    type = "timestamp"
    resolver "ISODateResolver" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ISODateResolver"
      path_resolver = true
    }
  }

  column "delete_time" {
    type = "timestamp"
    resolver "ISODateResolver" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ISODateResolver"
      path_resolver = true
    }
  }

  column "update_time" {
    type = "timestamp"
    resolver "ISODateResolver" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ISODateResolver"
      path_resolver = true
    }
  }

  ##################################################################
  # Skip columns not currently supported by Cloud Run
  ##################################################################
  column "metadata_cluster_name" {
    skip = true
  }
  column "metadata_deletion_grace_period_seconds" {
    skip = true
  }
  column "metadata_deletion_timestamp" {
    skip = true
  }
  column "metadata_finalizers" {
    skip = true
  }
  column "metadata_generate_name" {
    skip = true
  }
  column "spec_template_metadata_cluster_name" {
    skip = true
  }
  column "spec_template_metadata_deletion_grace_period_seconds" {
    skip = true
  }
  column "spec_template_metadata_deletion_timestamp" {
    skip = true
  }
  column "spec_template_metadata_finalizers" {
    skip = true
  }
  column "spec_template_metadata_generate_name" {
    skip = true
  }

  ##################################################################
  # Relations
  ##################################################################
  relation "gcp" "cloudrun" "metadata_owner_references" {
    ignore_in_tests = true
  }

  relation "gcp" "cloudrun" "spec_template_metadata_owner_references" {
    ignore_in_tests = true
  }

  relation "gcp" "cloudrun" "spec_template_spec_volumes" {
    path = "google.golang.org/api/run/v1.Volume"
    description = "Volume represents a named volume in a container"
    rename = "spec_template_volumes"

    relation "gcp" "cloudrun" "config_map_items" {
      ignore_in_tests = true
    }
  }

  relation "gcp" "cloudrun" "spec_template_spec_containers" {
    path = "google.golang.org/api/run/v1.Container"
    description = "A single application container"
    rename = "spec_template_containers"

    ignore_columns_in_tests = [
      "readiness_probe_http_get_http_headers",
      "readiness_probe_exec_command",
      "startup_probe_exec_command",
      "args",
      "resources_requests",
      "liveness_probe_exec_command",
    ]

    column "ports" {
      type = "json"
      generate_resolver = false
    }

    column "readiness_probe_http_get_http_headers" {
      type = "json"
      generate_resolver = false
    }

    ##################################################################
    # Skip columns not currently supported by Cloud Run
    ##################################################################
    column "env_from" {
      skip = true
    }
    column "liveness_probe_http_get_http_headers" {
      skip = true
    }
    column "startup_probe_http_get_http_headers" {
      skip = true
    }
    column "readiness_probe_http_get_http_headers" {
      skip = true
    }
  }
}
