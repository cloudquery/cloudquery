service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}


resource "gcp" "serviceusage" "services" {
  path = "google.golang.org/api/serviceusage/v1.GoogleApiServiceusageV1Service"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreErrorHandler"
  }
  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/plugins/source/gcp/client.ProjectMultiplex"
  }
  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/plugins/source/gcp/client.DeleteProjectFilter"
  }

  userDefinedColumn "project_id" {
    type        = "string"
    description = "GCP Project Id of the resource"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/plugins/source/gcp/client.ResolveProject"
    }
  }
  options {
    primary_keys = ["name"]
  }
  column "name" {
    skip = true
  }
  userDefinedColumn "name" {
    type = "string"
    resolver "pathResolver" {
      path   = "github.com/cloudquery/plugin-sdk/schema.PathResolver"
      params = ["Config.Name"]
    }
  }
  column "config" {
    skip_prefix = true
  }
  column "documentation" {
    type              = "json"
    generate_resolver = true
  }

  column "authentication" {
    type              = "json"
    generate_resolver = true
  }
  relation "gcp" "serviceusage" "monitored_resources" {
    column "labels" {
      type              = "json"
      generate_resolver = true
    }
  }

  relation "gcp" "serviceusage" "apis" {
    column "methods" {
      type              = "json"
      generate_resolver = true
    }

    column "mixins" {
      type              = "json"
      generate_resolver = true
    }
    column "options" {
      type              = "json"
      generate_resolver = true
    }
  }

  relation "gcp" "serviceusage" "quota_limits" {
    column "default_limit" {
      type              = "int"
      generate_resolver = true
    }
  }
}

