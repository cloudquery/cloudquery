data "google_compute_image" "gcp_compute_images_image" {
  name    = "cos-89-16108-470-16"
  project = "cos-cloud"
}

resource "google_compute_instance" "gcp_compute_images_instance" {
  name         = "gcp-compute-images-compute-instance-${var.test_suffix}"
  machine_type = "f1-micro"
  zone         = "${var.region}-a"

  tags = [
  "test"]


  boot_disk {
    initialize_params {
      image = data.google_compute_image.gcp_compute_images_image.self_link
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
    email = google_service_account.service_account.email
    scopes = [
    "cloud-platform"]
  }
}