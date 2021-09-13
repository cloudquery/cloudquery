resource "google_compute_address" "compute_address_1" {
  name         = "compute-addr-${var.test_prefix}${var.test_suffix}"
  description  = "my description"
  subnetwork   = google_compute_subnetwork.network-subnetwork.id
  address_type = "INTERNAL"
  address      = "10.2.133.133"
}
