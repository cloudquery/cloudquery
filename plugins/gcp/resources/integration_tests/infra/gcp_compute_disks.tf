resource "google_compute_disk" "gcp_compute_disks_disk" {
  name  = "gcp-compute-disks-disk-${var.test_suffix}"
  type  = "pd-ssd"
  zone  = "${var.region}-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  size                      = 10
  physical_block_size_bytes = 4096
}