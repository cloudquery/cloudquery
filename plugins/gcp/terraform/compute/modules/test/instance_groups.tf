resource "google_compute_instance_group" "test" {
  name        = "${var.prefix}-instance-group"
  description = "Integration test instance group"
  zone        = "${var.region}-a"
  network     = module.vpc.network_id
}