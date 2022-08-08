resource "google_compute_instance_group" "test" {
  name        = "${var.prefix}-instance-group"
  description = "Integration test instance group"
  zone        = data.google_compute_zones.available.names[0]
  network     = module.vpc.network_id
}

data "google_compute_zones" "available" {
}