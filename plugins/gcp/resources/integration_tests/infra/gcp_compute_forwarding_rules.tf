resource "google_compute_region_backend_service" "gcp_forwarding_rules_backend_svc" {
  name                  = "backend-svc-${var.test_prefix}${var.test_suffix}"
  provider              = google-beta
  region                = var.region
  load_balancing_scheme = "EXTERNAL"
  health_checks = [
  google_compute_region_health_check.gcp_forwarding_rules_health_check.id]

  backend {
    group          = google_compute_region_instance_group_manager.gcp_forwarding_rules_instance_group_manager.instance_group
    balancing_mode = "CONNECTION"
  }
}

# MIG
resource "google_compute_region_instance_group_manager" "gcp_forwarding_rules_instance_group_manager" {
  name     = "forwarding-rules-igm-${var.test_prefix}${var.test_suffix}"
  provider = google-beta
  region   = var.region
  version {
    instance_template = google_compute_instance_template.gcp_forwarding_rules_instance_template.id
    name              = "primary"
  }
  base_instance_name = "vm"
  target_size        = 1
}


# instance template
resource "google_compute_instance_template" "gcp_forwarding_rules_instance_template" {
  name         = "forwarding-rules-it-${var.test_prefix}${var.test_suffix}"
  provider     = google-beta
  machine_type = "e2-small"
  tags = [
    "allow-ssh",
  "allow-health-check"]

  network_interface {
    network    = google_compute_network.network.id
    subnetwork = google_compute_subnetwork.network-subnetwork.id
    access_config {
      # add external ip to fetch packages
    }
  }
  disk {
    source_image = "debian-cloud/debian-10"
    auto_delete  = true
    boot         = true
  }

  # install nginx and serve a simple web page
  metadata = {
    startup-script = <<-EOF1
      #! /bin/bash
      set -euo pipefail

      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y nginx-light jq

      NAME=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/hostname")
      IP=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip")
      METADATA=$(curl -f -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=True" | jq 'del(.["startup-script"])')

      cat <<EOF > /var/www/html/index.html
      <pre>
      Name: $NAME
      IP: $IP
      Metadata: $METADATA
      </pre>
      EOF
    EOF1
  }
  lifecycle {
    create_before_destroy = true
  }
}

# health check
resource "google_compute_region_health_check" "gcp_forwarding_rules_health_check" {
  name               = "health-check-${var.test_prefix}${var.test_suffix}"
  provider           = google-beta
  check_interval_sec = 1
  timeout_sec        = 1
  region             = var.region

  tcp_health_check {
    port = "80"
  }
}


# forwarding rule
resource "google_compute_forwarding_rule" "google_compute_forwarding_rule" {
  name            = "forwarding-rule-${var.test_prefix}${var.test_suffix}"
  provider        = google-beta
  region          = var.region
  port_range      = 80
  backend_service = google_compute_region_backend_service.gcp_forwarding_rules_backend_svc.id

  labels = {
    "test" = "test"
  }
}