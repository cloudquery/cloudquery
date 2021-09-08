resource "digitalocean_project" "do_project" {
  name        = "do_project-${var.test_prefix}-${var.test_suffix}"
  description = "A project for e2e testing"
  purpose     = "E2E"
  environment = "Development"

  resources = [digitalocean_droplet.do_droplet_other_project.urn]
}
