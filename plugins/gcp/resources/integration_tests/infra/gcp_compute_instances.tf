resource "google_compute_instance" "google_compute_instances_instance" {
  name         = "compute-instance-${var.test_prefix}${var.test_suffix}"
  machine_type = "f1-micro"
  zone         = "${var.region}-a"

  tags = [
  "test"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral IP
    }
  }

  metadata = {
    test = "test"
  }

  metadata_startup_script = "echo hi > /test.txt"
  service_account {
    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    email  = google_service_account.service_account.email
    scopes = ["cloud-platform"]
  }

}

