resource "digitalocean_vpc" "do_vpc" {
  name     = "dovpc${var.test_suffix}"
  region   = "nyc3"
  ip_range = "10.10.10.0/24"
}