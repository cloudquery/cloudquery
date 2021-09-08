resource "digitalocean_volume" "do_volume" {
  region                  = "nyc3"
  name                    = "dovolume${var.test_suffix}"
  size                    = 100
  initial_filesystem_type = "ext4"
  description             = "test volume"
}

resource "digitalocean_volume_attachment" "foobar" {
  droplet_id = digitalocean_droplet.do_droplet.id
  volume_id  = digitalocean_volume.do_volume.id
}