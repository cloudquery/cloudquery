resource "digitalocean_container_registry" "do_cr" {
  name                   = random_id.test_id.hex
  subscription_tier_slug = "starter"
}