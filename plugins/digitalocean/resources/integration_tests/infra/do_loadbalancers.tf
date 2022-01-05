resource "digitalocean_loadbalancer" "do_loadbalancer" {
  name   = "do-loadbalancer${random_id.test_id.hex}"
  region = "nyc3"

  forwarding_rule {
    entry_port     = 1245
    entry_protocol = "http"

    target_port     = 3030
    target_protocol = "http"
  }

  healthcheck {
    port     = 22
    protocol = "tcp"
  }

  droplet_ids = [digitalocean_droplet.do_droplet_other_project.id]
}