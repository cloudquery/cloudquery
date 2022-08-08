service          = "azure"
output_directory = "."
add_generate     = true

resource "azure" "" "front_doors" {
  path        = "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor.FrontDoor"
  description = "Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there."

  userDefinedColumn "subscription_id" {
    type        = "string"
    description = "Azure subscription ID"
    resolver "resolveAzureSubscription" {
      path = "github.com/cloudquery/cq-provider-azure/client.ResolveAzureSubscription"
    }
  }

  options {
    primary_keys = [
      "subscription_id",
      "id"
    ]
  }

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.DeleteSubscriptionFilter"
  }

  column "id" {
    description = "Resource ID"
  }

  column "name" {
    description = "Resource name"
  }

  column "type" {
    description = "Resource type"
  }

  column "location" {
    description = "Resource location"
  }

  column "tags" {
    description = "Resource tags"
  }

  column "properties_resource_state" {
    rename      = "resource_state"
    description = "Resource state of the Front Door"
  }

  column "properties_provisioning_state" {
    rename      = "provisioning_state"
    description = "Provisioning state of the Front Door"
  }

  column "properties_cname" {
    rename      = "cname"
    description = "The host that each frontend endpoint must CNAME to"
  }

  column "properties_frontdoor_id" {
    rename      = "frontdoor_id"
    description = "The ID of the Front Door"
  }

  column "properties_friendly_name" {
    rename      = "friendly_name"
    description = "A friendly name for the Front Door"
  }

  column "properties_backend_pools_settings_enforce_certificate_name_check" {
    rename      = "enforce_certificate_name_check"
    description = "Whether to enforce certificate name check on HTTPS requests to all backend pools"
  }

  column "properties_backend_pools_settings_send_recv_timeout_seconds" {
    rename      = "send_recv_timeout_seconds"
    description = "Send and receive timeout on forwarding request to the backend"
  }

  column "properties_enabled_state" {
    rename      = "enabled_state"
    description = "Operational status of the Front Door load balancer"
  }

  relation "azure" "front_door" "properties_rules_engines" {
    rename          = "rules_engines"
    description     = "Rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response."
    ignore_in_tests = true

    column "rules_engine_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource state"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }

    relation "azure" "front_door" "rules_engine_properties_rules" {
      rename      = "rules"
      description = "A list of rules that define a particular Rules Engine Configuration."

      column "name" {
        description = "A name to refer to the rule"
      }

      column "priority" {
        description = "A priority assigned to the rule"
      }

      column "match_processing_behavior" {
        description = "If the rule is a match should the rules engine continue running the remaining rules or stop"
      }

      column "action_route_configuration_override" {
        rename            = "route_configuration_override"
        description       = "Override the route configuration"
        type              = "json"
        generate_resolver = true
      }

      column "action_request_header_actions" {
        rename            = "request_header_actions"
        description       = "A list of header actions to apply from the request from AFD to the origin."
        type              = "json"
        generate_resolver = true
      }
      column "action_response_header_actions" {
        rename            = "response_header_actions"
        description       = "A list of header actions to apply from the response from AFD to the client."
        type              = "json"
        generate_resolver = true
      }

      relation "azure" "front_door" "match_conditions" {
        description = "A list of match conditions that must meet in order for the actions of the rule to run. Having no match conditions means the actions will always run."

        column "rules_engine_match_variable" {
          rename      = "match_variable"
          description = "Match variable"
        }

        column "selector" {
          description = "Name of selector in request header or request body to be matched"
        }

        column "rules_engine_operator" {
          rename      = "operator"
          description = "Describes operator to apply to the match condition"
        }

        column "negate_condition" {
          description = "Describes if this is negate condition or not"
        }

        column "rules_engine_match_value" {
          rename      = "match_value"
          description = "Match values to match against"
        }

        column "transforms" {
          description = "List of transforms"
        }
      }
    }
  }

  relation "azure" "frontdoor" "properties_routing_rules" {
    rename      = "routing_rules"
    description = "Routing rules represent specifications for traffic to treat and where to send it, along with health probe information."

    ignore_columns_in_tests = [
      "routing_rule_properties_rules_engine_id",
      "routing_rule_properties_web_application_firewall_policy_link_id",
    ]

    column "routing_rule_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource state"
    }

    column "routing_rule_properties_accepted_protocols" {
      rename      = "accepted_protocols"
      description = "Protocol schemes to match for the rule"
    }

    column "routing_rule_properties_patterns_to_match" {
      rename      = "patterns_to_match"
      description = "The route patterns of the rule"
    }

    column "routing_rule_properties_enabled_state" {
      rename      = "enabled_state"
      description = "Whether the rule is enabled"
    }

    column "routing_rule_properties_route_configuration" {
      rename            = "route_configuration"
      description       = "A reference to the routing configuration"
      type              = "json"
      generate_resolver = true
    }

    column "routing_rule_properties_rules_engine_id" {
      rename      = "rules_engine_id"
      description = "ID of a specific Rules Engine Configuration to apply to the route"
    }

    column "routing_rule_properties_web_application_firewall_policy_link_id" {
      rename      = "web_application_firewall_policy_link_id"
      description = "ID of the Web Application Firewall policy for each routing rule (if applicable)"
    }

    column "routing_rule_properties_frontend_endpoints" {
      rename            = "frontend_endpoints"
      description       = "Frontend endpoints associated with the rule"
      type              = "StringArray"
      generate_resolver = true
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }

  }

  relation "azure" "frontdoor" "properties_load_balancing_settings" {
    rename      = "load_balancing_settings"
    description = "Load balancing settings for a backend pool associated with the Front Door instance"

    column "load_balancing_settings_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource state"
    }

    column "load_balancing_settings_properties_sample_size" {
      rename      = "sample_size"
      description = "The number of samples to consider for load balancing decisions"
    }

    column "load_balancing_settings_properties_successful_samples_required" {
      rename      = "successful_samples_required"
      description = "The number of samples within the sample period that must succeed"
    }

    column "load_balancing_settings_properties_additional_latency_milliseconds" {
      rename      = "additional_latency_milliseconds"
      description = "The additional latency in milliseconds for probes to fall into the lowest latency bucket"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }
  }

  relation "azure" "frontdoor" "properties_health_probe_settings" {
    rename      = "health_probe_settings"
    description = "Health probe settings for a backend pool associated with this Front Door instance"

    column "health_probe_settings_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource state"
    }

    column "health_probe_settings_properties_path" {
      rename      = "path"
      description = "The path to use for the health probe"
    }

    column "health_probe_settings_properties_protocol" {
      rename      = "protocol"
      description = "Protocol scheme to use for the health probe"
    }

    column "health_probe_settings_properties_interval_in_seconds" {
      rename      = "interval_in_seconds"
      description = "The number of seconds between health probes"
    }

    column "health_probe_settings_properties_health_probe_method" {
      rename      = "health_probe_method"
      description = "Which HTTP method is used to perform the health probe"
    }

    column "health_probe_settings_properties_enabled_state" {
      rename      = "enabled_state"
      description = "Whether the health probe is enabled"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }
  }

  relation "azure" "frontdoor" "properties_backend_pools" {
    rename      = "backend_pools"
    description = "Backend pools available to routing rules"

    column "backend_pool_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource state"
    }

    column "backend_pool_properties_load_balancing_settings_id" {
      rename      = "load_balancing_settings_id"
      description = "Load balancing settings ID for the backend pool"
    }

    column "backend_pool_properties_health_probe_settings_id" {
      rename      = "health_probe_settings_id"
      description = "L7 health probe settings ID for the backend pool"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }

    relation "azure" "frontdoor" "backend_pool_properties_backends" {
      rename      = "backends"
      description = "The set of backends for the backend pool"

      ignore_columns_in_tests = [
        "private_link_alias",
        "private_link_resource_id",
        "private_link_location",
        "private_link_approval_message",
      ]

      column "address" {
        description = "Location of the backend (IP address or FQDN)"
      }

      column "private_link_alias" {
        description = "The Alias of the Private Link resource"
      }

      column "private_link_resource_id" {
        description = "The Resource ID of the Private Link resource"
      }

      column "private_link_location" {
        description = "The location of the Private Link resource"
      }

      column "private_endpoint_status" {
        description = "The Approval status for the connection to the Private Link"
      }

      column "private_link_approval_message" {
        description = "A custom message to be included in the approval request to connect to the Private Link"
      }

      column "http_port" {
        description = "The HTTP TCP port number"
      }

      column "https_port" {
        description = "The HTTPS TCP port number"
      }

      column "enabled_state" {
        description = "Whether the use of the backend is enabled"
      }

      column "priority" {
        description = "Priority to use for load balancing"
      }

      column "weight" {
        description = "Weight of the endpoint for load balancing purposes"
      }

      column "backend_host_header" {
        rename      = "host_header"
        description = "The value to use as the host header sent to the backend"
      }
    }
  }

  relation "azure" "frontdoor" "properties_frontend_endpoints" {
    rename      = "frontend_endpoints"
    description = "Frontend endpoints available to routing rules"

    ignore_columns_in_tests = [
      "frontend_endpoint_properties_custom_https_configuration_protocol_type",
      "frontend_endpoint_properties_id",
      "frontend_endpoint_properties_custom_https_configuration_vault_id",
      "frontend_endpoint_properties_custom_https_configuration_secret_name",
      "frontend_endpoint_properties_custom_https_configuration_secret_version",
    ]

    column "frontend_endpoint_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource state"
    }

    column "frontend_endpoint_properties_custom_https_provisioning_state" {
      rename      = "custom_https_provisioning_state"
      description = "Provisioning status of custom https of the frontend endpoint"
    }

    column "frontend_endpoint_properties_custom_https_provisioning_substate" {
      rename      = "custom_https_provisioning_substate"
      description = "Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step"
    }

    column "frontend_endpoint_properties_custom_https_configuration_certificate_source" {
      rename      = "certificate_source"
      description = "Defines the source of the SSL certificate"
    }

    column "frontend_endpoint_properties_custom_https_configuration_protocol_type" {
      rename      = "protocol_type"
      description = "Defines the TLS extension protocol that is used for secure delivery"
    }

    column "frontend_endpoint_properties_custom_https_configuration_minimum_tls_version" {
      rename      = "minimum_tls_version"
      description = "The minimum TLS version required from the clients to establish an SSL handshake with Front Door"
    }

    column "frontend_endpoint_properties_custom_https_configuration_key_vault_certificate_source_parameters" {
      rename      = "key_vault_certificate_source_parameters"
      skip_prefix = true
    }
    column "frontend_endpoint_properties_custom_https_configuration_vault_id" {
      rename      = "vault_id"
      description = "The Key Vault containing the SSL certificate"
    }

    column "frontend_endpoint_properties_custom_https_configuration_secret_name" {
      rename      = "secret_name"
      description = "The name of the Key Vault secret representing the full certificate PFX"
    }

    column "frontend_endpoint_properties_custom_https_configuration_secret_version" {
      rename      = "secret_version"
      description = "The version of the Key Vault secret representing the full certificate PFX"
    }

    column "frontend_endpoint_properties_custom_https_configuration_certificate_source_parameters" {
      rename      = "certificate_source_parameters"
      skip_prefix = true
    }

    column "frontend_endpoint_properties_custom_https_configuration_certificate_type" {
      rename      = "certificate_type"
      description = "The type of the certificate used for secure connections to the frontend endpoint"
    }

    column "frontend_endpoint_properties_host_name" {
      rename      = "host_name"
      description = "The host name of the frontend endpoint"
    }

    column "frontend_endpoint_properties_session_affinity_enabled_state" {
      rename      = "session_affinity_enabled_state"
      description = "Whether session affinity is allowed on the host"
    }

    column "frontend_endpoint_properties_session_affinity_ttl_seconds" {
      rename      = "session_affinity_ttl_seconds"
      description = "The TTL to use in seconds for session affinity, if applicable"
    }

    column "frontend_endpoint_properties_web_application_firewall_policy_link" {
      rename      = "web_application_firewall_policy_link"
      skip_prefix = true
    }

    column "frontend_endpoint_properties_id" {
      rename      = "web_application_firewall_policy_link_id"
      description = "Defines the Web Application Firewall policy for each host (if applicable)"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }
  }
}