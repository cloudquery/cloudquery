resource "google_compute_autoscaler" "gcp_compute_autoscalers_autoscaler" {
  name   = "autoscaler${var.test_prefix}${var.test_suffix}"
  zone   = "${var.region}-f"
  target = google_compute_instance_group_manager.gcp_compute_autoscalers_instance_group_manager.id

  autoscaling_policy {
    max_replicas    = 5
    min_replicas    = 1
    cooldown_period = 60

    cpu_utilization {
      target = 0.5
    }

    metric {
      name = "pubsub.googleapis.com/subscription/num_undelivered_messages"
      //      filter = "resource.type = pubsub_subscription AND resource.label.subscription_id = \"test\""
      target = 100
      type   = "DELTA_PER_MINUTE"
    }
  }


}

resource "google_compute_instance_template" "gcp_compute_autoscalers_instance_template" {
  name           = "autoscalers-it-${var.test_prefix}${var.test_suffix}"
  machine_type   = "f1-micro"
  can_ip_forward = false

  tags = [
    "foo",
  "bar"]

  disk {
    source_image = data.google_compute_image.debian_9.id
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  service_account {
    scopes = [
      "userinfo-email",
      "compute-ro",
    "storage-ro"]
  }
}

resource "google_compute_target_pool" "gcp_compute_autoscalers_target_pool" {
  name = "autoscalers-tp-${var.test_prefix}${var.test_suffix}"
}

resource "google_compute_instance_group_manager" "gcp_compute_autoscalers_instance_group_manager" {
  name = "autoscalers-igm-${var.test_prefix}${var.test_suffix}"
  zone = "${var.region}-f"

  version {
    instance_template = google_compute_instance_template.gcp_compute_autoscalers_instance_template.id
    name              = "primary"
  }

  target_pools = [
  google_compute_target_pool.gcp_compute_autoscalers_target_pool.id]
  base_instance_name = "foobar"
}

data "google_compute_image" "debian_9" {
  family  = "debian-9"
  project = "debian-cloud"
}