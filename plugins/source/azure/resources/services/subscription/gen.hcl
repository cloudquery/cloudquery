service          = "azure"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["READ-ONLY; "]
}

description_modifier "remove_field_name" {
  regex = ".+- "
}

resource "azure" "subscription" "subscriptions" {
  path = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions.Subscription"
  description = "Azure subscription information"

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

  column "subscription_policies_location_placement_id" {
    rename = "location_placement_id"
  }

  column "subscription_policies_quota_id" {
    rename = "quota_id"
  }

  column "subscription_policies_spending_limit" {
    rename = "spending_limit"
  }

  column "managed_by_tenants" {
    type = "stringArray"
    generate_resolver = true
  }

  ignore_columns_in_tests = ["tags"]

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.DeleteSubscriptionFilter"
  }
}

resource "azure" "subscription" "tenants" {
  path = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions.TenantIDDescription"
  description = "Azure tenant information"

  options {
    primary_keys = [
      "id"
    ]
  }

  column "id" {
    type = "string"
    description = "The fully qualified ID of the tenant"
  }

  ignore_columns_in_tests = ["country","country_code","display_name","tenant_branding_logo_url","default_domain","domains","tenant_type"]

  multiplex "SingleSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }
}
