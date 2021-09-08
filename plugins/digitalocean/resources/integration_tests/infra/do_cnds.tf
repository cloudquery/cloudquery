resource "digitalocean_cdn" "do_cdn" {
  origin = digitalocean_spaces_bucket.do_spaces.bucket_domain_name

  depends_on = [digitalocean_spaces_bucket.do_spaces]
}