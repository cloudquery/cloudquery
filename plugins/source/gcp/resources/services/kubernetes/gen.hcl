service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}


resource "gcp" "kubernetes" "clusters" {
  path = "google.golang.org/api/container/v1.Cluster"
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

  column "cluster_ipv4_cidr" {
    type = "cidr"
  }

  column "ip_allocation_policy_cluster_ipv4_cidr" {
    type = "cidr"
  }

  column "ip_allocation_policy_cluster_ipv4_cidr_block" {
    type = "cidr"
  }

  column "ip_allocation_policy_node_ipv4_cidr" {
    type = "cidr"
  }

  column "ip_allocation_policy_node_ipv4_cidr_block" {
    type = "cidr"
  }

  column "ip_allocation_policy_services_ipv4_cidr_block" {
    type = "cidr"
  }

  column "ip_allocation_policy_tpu_ipv4_cidr_block" {
    type = "cidr"
  }

  column "private_cluster_config_master_ipv4_cidr_block" {
    type = "cidr"
  }

  column "services_ipv4_cidr" {
    type = "cidr"
  }

  column "tpu_ipv4_cidr_block" {
    type = "cidr"
  }

  column "ip_allocation_policy_services_ipv4_cidr" {
    type = "cidr"
  }

  column "create_time" {
    type = "timestamp"
  }

  column "expire_time" {
    type = "timestamp"
  }

  column "endpoint" {
    type = "inet"
  }

  column "maintenance_policy_window_daily_maintenance_window_start_time" {
    type = "timestamp"
  }

  column "maintenance_policy_window_daily_maintenance_window_end_time" {
    type = "timestamp"
  }

  column "maintenance_policy_window_recurring_window_window_start_time" {
    type = "timestamp"
  }

  column "maintenance_policy_window_recurring_window_window_end_time" {
    type = "timestamp"
  }

  column "autoscaling_autoprovisioning_node_pool_defaults" {
    type = "json"
  }

  column "autoscaling_resource_limits" {
    type = "json"
  }

  column "master_authorized_networks_config_cidr_blocks" {
    type = "json"
  }

  column "conditions" {
    type = "json"
  }

  column "node_config" {
    type = "json"
  }

  column "network_config_pod_ipv4_cidr_block" {
    type = "cidr"
  }

  column "management_upgrade_options_auto_upgrade_start_time" {
    type = "timestamp"
  }

  column "resource_usage_export_config" {
    type = "json"
  }

  relation "gcp" "kubernetes" "node_pools" {
    column "config_accelerators" {
      type = "json"
    }

    column "conditions" {
      type = "json"
    }

    column "config_taints" {
      type = "json"
    }

    column "management_upgrade_options_auto_upgrade_start_time" {
      type = "timestamp"
    }

    column "network_config_pod_ipv4_cidr_block" {
      type = "cidr"
    }
  }

  options {
    primary_keys = ["id"]
  }
}

