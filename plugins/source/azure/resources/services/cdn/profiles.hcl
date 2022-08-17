service          = "azure"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["READ-ONLY; "]
}

description_modifier "remove_field_name" {
  regex = ".+- "
}

resource "azure" "cdn" "profiles" {
  path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.Profile"
  userDefinedColumn "subscription_id" {
    type        = "string"
    description = "Azure subscription ID"
    resolver "resolveAzureSubscription" {
      path = "github.com/cloudquery/cloudquery/plugins/source/azure/client.ResolveAzureSubscription"
    }
  }

  options {
    primary_keys = [
      "subscription_id",
      "id"
    ]
  }

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cloudquery/plugins/source/azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cloudquery/plugins/source/azure/client.DeleteSubscriptionFilter"
  }

  column "profile_properties" {
    skip_prefix = true
  }

  column "system_data" {
    skip_prefix = true
  }

  column "resource_state" {
    rename = "state"
  }


  user_relation "azure" "cdn" "endpoints" {
    path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.Endpoint"
    column "endpoint_properties" {
      skip_prefix = true
    }
    column "system_data" {
      skip_prefix = true
    }

    //todo remove in fixed sdk
    relation "azure" "cdn" "url_signing_keys" {
      resolver "resolver" {
        generate = true
      }
    }
    relation "azure" "cdn" "geo_filters" {
      resolver "resolver" {
        generate = true
      }
    }

    relation "azure" "cdn" "origin_groups" {
      resolver "resolver" {
        generate = true
      }
      column "deep_created_origin_group_properties" {
        skip_prefix = true
      }
      column "response_based_origin_error_detection_settings" {
        type = "json"
      }

      column "origins" {
        type              = "stringarray"
        generate_resolver = true
      }
    }

    relation "azure" "cdn" "origins" {
      resolver "resolver" {
        generate = true
      }
      column "deep_created_origin_properties" {
        skip_prefix = true
      }
    }


    relation "azure" "cdn" "delivery_policy_rules" {
      //todo check type
      column "conditions" {
        type              = "json"
        generate_resolver = true
      }
      //todo check type
      column "actions" {
        type              = "json"
        generate_resolver = true
      }
      resolver "resolver" {
        generate = true
      }
    }
    user_relation "azure" "cdn" "custom_domains" {
      path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.CustomDomain"

      column "custom_domain_properties" {
        skip_prefix = true
      }
      //todo check type
      column "custom_https_parameters" {
        type = "json"
      }
      column "resource_state" {
        rename = "state"
      }
      column "system_data" {
        skip_prefix = true
      }
    }

    user_relation "azure" "cdn" "routes" {
      path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.Route"
      column "route_properties" {
        skip_prefix = true
      }

      column "compression_settings" {
        type              = "json"
        generate_resolver = true
      }
      column "system_data" {
        skip_prefix = true
      }

      column "rule_sets" {
        type              = "stringarray"
        generate_resolver = true
      }

      column "custom_domains" {
        type              = "stringarray"
        generate_resolver = true
      }
    }
  }

  user_relation "azure" "cdn" "rule_sets" {
    path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.RuleSet"

    column "rule_set_properties" {
      skip_prefix = true
    }

    column "system_data" {
      skip_prefix = true
    }

    user_relation "azure" "cdn" "rules" {
      path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.Rule"

      column "rule_properties" {
        skip_prefix = true
      }

      column "conditions" {
        type              = "json"
        generate_resolver = true
      }

      column "actions" {
        type = "json"
        generate_resolver = true
      }

      column "system_data" {
        skip_prefix = true
      }
    }
  }

  user_relation "azure" "cdn" "security_policies" {
    path = "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn.SecurityPolicy"
    column "security_policy_properties" {
      skip_prefix = true
    }

    column "parameters" {
      #      skip = true
      type = "json"
    }

    #    userDefinedColumn "WebApplicationFirewallParameters" {
    #      type              = "json"
    #      generate_resolver = true
    #    }
    #
    #    userDefinedColumn "web_application_firewall_parameters" {
    #      type              = "json"
    #      generate_resolver = true
    #    }

    #    userDefinedColumn "parameters_type" {
    #      type              = "string"
    #      generate_resolver = true
    #    }
    column "system_data" {
      skip_prefix = true
    }
  }
}