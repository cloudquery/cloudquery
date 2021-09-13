resource "google_compute_url_map" "urlmaps_urlmap" {
  name        = "urlmap${var.test_prefix}${var.test_suffix}"
  description = "a description"

  //  region = var.region
  default_service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id

  host_rule {
    hosts = [
    "mysite.com"]
    path_matcher = "mysite"
  }

  host_rule {
    hosts = [
    "myothersite.com"]
    path_matcher = "otherpaths"
  }

  path_matcher {
    name            = "mysite"
    default_service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id

    path_rule {
      paths = [
      "/home"]
      service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
    }

    path_rule {
      paths = [
      "/login"]
      service = google_compute_backend_service.urlmaps_storage_backend_svc.id
    }

    path_rule {
      paths = [
      "/static"]
      service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
    }
  }

  path_matcher {
    name            = "otherpaths"
    default_service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
  }

  test {
    service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
    host    = "hi.com"
    path    = "/home"
  }

  depends_on = [
  google_compute_backend_bucket.urlmaps_compute_backend_bucket]
}


resource "google_compute_backend_bucket" "urlmaps_compute_backend_bucket" {
  name        = "static-asset-backend-${var.test_prefix}${var.test_suffix}"
  bucket_name = google_storage_bucket.urlmaps_storage_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "urlmaps_storage_bucket" {
  name     = "static-asset-bucket-${var.test_prefix}${var.test_suffix}"
  location = "US"
}

resource "google_compute_backend_service" "urlmaps_storage_backend_svc" {
  name     = "url-maps-backend-svc-${var.test_prefix}${var.test_suffix}"
  provider = google-beta
  //  region = var.region
  load_balancing_scheme = "EXTERNAL"
  health_checks = [
  google_compute_health_check.urlmaps_health_check.id]

  backend {
    group          = google_compute_region_instance_group_manager.urlmaps_instance_group_manager.instance_group
    balancing_mode = "UTILIZATION"
  }
}


# MIG
resource "google_compute_region_instance_group_manager" "urlmaps_instance_group_manager" {
  name     = "url-maps-igm-${var.test_prefix}${var.test_suffix}"
  provider = google-beta
  region   = var.region
  version {
    instance_template = google_compute_instance_template.urlmaps_instance_template.id
    name              = "primary"
  }
  base_instance_name = "vm"
  target_size        = 1
}


# instance template
resource "google_compute_instance_template" "urlmaps_instance_template" {
  name         = "url-maps-it-${var.test_prefix}${var.test_suffix}"
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
resource "google_compute_health_check" "urlmaps_health_check" {
  name               = "url-maps-health-check-${var.test_prefix}${var.test_suffix}"
  provider           = google-beta
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}