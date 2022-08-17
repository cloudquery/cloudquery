service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}


resource "gcp" "cloudbilling" "accounts" {
  path = "github.com/cloudquery/plugins/source/gcp/resources/services/cloudbilling.BillingAccountWrapper"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreErrorHandler"
  }

  options {
    primary_keys = [
      "project_id",
      "name"
    ]
  }
  
  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/plugins/source/gcp/client.ProjectMultiplex"
  }

  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/plugins/source/gcp/client.DeleteProjectFilter"
  }

  column "billing_account" {
    skip_prefix = true
  }


  column "name" {
    description = "The resource name of the billing account."
  }
  column "open" {
    description = "True if the billing account is open"
  }

  column "project_billing_info_billing_account_name" {
    skip = true
  }

  column "project_billing_info_billing_enabled" {
    rename = "project_billing_enabled"
  }

  column "project_billing_info_name" {
    rename = "project_name"
  }

  column "project_billing_info_project_id" {
    rename = "project_id"
  }
}


resource "gcp" "cloudbilling" "services" {
  path = "google.golang.org/api/cloudbilling/v1.Service"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/plugins/source/gcp/client.IgnoreErrorHandler"
  }

  user_relation "gcp" "cloudbilling" "skus" {
    path = "google.golang.org/api/cloudbilling/v1.Sku"

    column "category" {
      skip_prefix = true
    }

    relation "gcp" "cloudbilling" "pricing_info" {

      column "aggregation_info" {
        skip_prefix = true
      }

      column "pricing_expression" {
        skip_prefix = true
      }
    }
  }
}