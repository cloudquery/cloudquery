# google_client_config and kubernetes provider must be explicitly specified like the following.
data "google_client_config" "default" {}

provider "kubernetes" {
  host                   = "https://${module.gke.endpoint}"
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(module.gke.ca_certificate)
}

locals {
  cluster_type           = "${lower(var.prefix)}-example-gke-cluster"
  network_name           = "${local.cluster_type}-network"
  subnet_name            = "${local.cluster_type}-subnet"
  master_auth_subnetwork = "${local.cluster_type}-master-subnet"
  pods_range_name        = "${local.cluster_type}-ip-range-pods"
  svc_range_name         = "${local.cluster_type}-ip-range-svc"
  subnet_names           = [for subnet_self_link in module.gcp-network.subnets_self_links : split("/", subnet_self_link)[length(split("/", subnet_self_link)) - 1]]
}

module "gke" {
  source                     = "terraform-google-modules/kubernetes-engine/google"
  project_id                 = var.project_id
  name                       = local.cluster_type
  region                     = var.region
  zones                      = var.zones
  network                    = module.gcp-network.network_name
  subnetwork                 = local.subnet_names[index(module.gcp-network.subnets_names, local.subnet_name)]
  ip_range_pods              = local.pods_range_name
  ip_range_services          = local.svc_range_name
  add_cluster_firewall_rules = true
  firewall_inbound_ports     = ["9443", "15017"]
  http_load_balancing        = false
  network_policy             = false
  horizontal_pod_autoscaling = true
  filestore_csi_driver       = false

  node_pools = [
    {
      name               = "${lower(var.prefix)}-example-gke-pool-1"
      machine_type       = "e2-micro"
      min_count          = 1
      max_count          = 1
      local_ssd_count    = 0
      disk_size_gb       = 10
      disk_type          = "pd-standard"
      image_type         = "COS_CONTAINERD"
      service_account    = "${lower(var.prefix)}-example-gke-service-account@${var.project_id}.iam.gserviceaccount.com"
      auto_upgrade       = true
      initial_node_count = 1
      preemptible        = true
    },
    {
      name               = "default-node-pool"
      machine_type       = "e2-micro"
      node_locations     = "us-east1-b,us-east1-c"
      min_count          = 1
      max_count          = 2
      local_ssd_count    = 0
      disk_size_gb       = 10
      disk_type          = "pd-standard"
      image_type         = "COS_CONTAINERD"
      enable_gcfs        = false
      auto_repair        = true
      auto_upgrade       = true
      service_account    = "${lower(var.prefix)}-example-gke-service-account@${var.project_id}.iam.gserviceaccount.com"
      preemptible        = true
      initial_node_count = 1
    },
  ]

  node_pools_oauth_scopes = {
    all = []

    default-node-pool = [
      "https://www.googleapis.com/auth/cloud-platform",
    ]
  }

  node_pools_labels = {
    all = {}

    default-node-pool = {
      default-node-pool = true
    }
  }

  node_pools_metadata = {
    all = {}

    default-node-pool = {
      node-pool-metadata-custom-value = "my-node-pool"
    }
  }

  node_pools_taints = {
    all = []

    default-node-pool = [
      {
        key    = "default-node-pool"
        value  = true
        effect = "PREFER_NO_SCHEDULE"
      },
    ]
  }

  node_pools_tags = {
    all = []

    default-node-pool = [
      "default-node-pool",
    ]
  }
}
