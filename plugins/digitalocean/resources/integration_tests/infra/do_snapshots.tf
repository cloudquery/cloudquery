resource "digitalocean_droplet_snapshot" "do_snapshot" {
  droplet_id = digitalocean_droplet.do_droplet.id
  name       = "do_image_snap${var.test_prefix}-${var.test_suffix}"

  depends_on = [digitalocean_droplet.do_droplet]
}