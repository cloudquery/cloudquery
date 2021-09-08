resource "digitalocean_container_registry" "do_cr" {
  name                   = "do-cr${var.test_prefix}-${var.test_suffix}"
  subscription_tier_slug = "starter"
}