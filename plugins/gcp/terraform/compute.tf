################################################################################
# Compute Module - Helpers
################################################################################

resource "google_compute_image" "image" {
  name = "${local.prefix}-compute-image"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }

  labels = local.labels
}

resource "google_compute_resource_policy" "regional-disk-replica-policy" {
  name = "${local.prefix}-regional-disk-replica-policy"
  region = local.region
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time = "04:00"
      }
    }
  }
}

resource "google_compute_region_disk_resource_policy_attachment" "regional-disk-replica-policy-attachment" {
  name = google_compute_resource_policy.regional-disk-replica-policy.name
  disk = google_compute_region_disk.region-disk.name
  region = local.region
}

resource "google_compute_region_disk" "region-disk" {
  name  = "${local.prefix}-disk"
  type  = "pd-ssd"
  region  = local.region
  labels = local.labels
  size = 10

  replica_zones = ["${local.region}-a", "${local.region}-b"]
}

################################################################################
# Compute Module
################################################################################

module "vm_instance_template" {
  source  = "terraform-google-modules/vm/google//modules/instance_template"
  version = "7.3.0"

  project_id      = local.project
  subnetwork      = module.vpc.subnets_ids.0
  service_account = {
    email  = module.service_accounts.email
    scopes = []
  }

  additional_networks = [
    {
      network = module.another-network.network_id
      subnetwork = module.another-network.subnets_ids.0
      subnetwork_project = null
      network_ip = ""
      access_config = []
    }
  ]

  disk_labels = local.labels

  name_prefix = local.prefix
  tags        = local.tags
  labels      = local.labels
}

module "compute_instance_group" {
  source  = "terraform-google-modules/vm/google//modules/mig"
  version = "7.3.0"

  project_id        = local.project
  region            = local.region
  hostname          = "${local.prefix}-group"
  instance_template = module.vm_instance_template.self_link
  target_size       = 1

  network = module.vpc.network_id
  #  target_pools              = 5
  #  distribution_policy_zones = var.distribution_policy_zones
  #  update_policy             = var.update_policy
  #  named_ports               = var.named_ports

  /* autoscaler */
  autoscaling_enabled = true
  max_replicas        = 5
  min_replicas        = 1
  cooldown_period     = 300
  autoscaling_metric  = [
    {
      name   = "pubsub.googleapis.com/subscription/num_undelivered_messages"
      type   = "GAUGE"
      target = 65535
    }
  ]
  #  autoscaling_cpu              = var.autoscaling_cpu
  #  autoscaling_metric           = var.autoscaling_metric
  #  autoscaling_lb               = var.autoscaling_lb
  #  autoscaling_scale_in_control = var.autoscaling_scale_in_control

}