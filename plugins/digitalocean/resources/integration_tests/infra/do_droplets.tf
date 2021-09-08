resource "digitalocean_droplet" "do_droplet" {
  image    = "ubuntu-18-04-x64"
  name     = "do-droplet${var.test_prefix}-${var.test_suffix}"
  region   = "nyc3"
  size     = "s-1vcpu-1gb"
  ipv6     = true
  vpc_uuid = digitalocean_vpc.do_vpc.id
}

resource "digitalocean_droplet" "do_droplet_other_project" {
  image  = "ubuntu-18-04-x64"
  name   = "do-droplet-other-project${var.test_prefix}-${var.test_suffix}"
  region = "nyc3"
  size   = "s-1vcpu-1gb"
  ipv6   = true
}