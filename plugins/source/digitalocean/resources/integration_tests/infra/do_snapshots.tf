resource "digitalocean_droplet_snapshot" "do_snapshot" {
  droplet_id = digitalocean_droplet.do_droplet.id
  name       = "do_image_snap${random_id.test_id.hex}"

  depends_on = [digitalocean_droplet.do_droplet]
}