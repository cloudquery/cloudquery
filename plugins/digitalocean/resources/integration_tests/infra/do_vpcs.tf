resource "digitalocean_vpc" "do_vpc" {
  name     = "dovpc${random_id.test_id.hex}"
  region   = "nyc3"
  ip_range = "10.10.10.0/24"
}