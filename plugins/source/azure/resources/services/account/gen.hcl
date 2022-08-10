service          = "azure"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["READ-ONLY; "]
}

description_modifier "remove_field_name" {
  regex = ".+- "
}

resource "azure" "account" "locations" {
  path = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions.Location"
  description = "Azure location information"

  userDefinedColumn "subscription_id" {
    type        = "string"
    description = "Azure subscription id"
    resolver "resolveAzureSubscription" {
      path = "github.com/cloudquery/cq-provider-azure/client.ResolveAzureSubscription"
    }
  }
  column "subscription_id" {
    skip = true
  }
  options {
    primary_keys = [
      "id"
    ]
  }

  column "metadata" {
    skip_prefix = true
  }

  ignore_columns_in_tests = ["home_location"]

  relation "azure" "account" "paired_region" {
    column "subscription_id" {
      skip = true
    }
  }

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.DeleteSubscriptionFilter"
  }
}