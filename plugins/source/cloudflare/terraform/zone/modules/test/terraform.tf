terraform {
  required_version = ">= 0.15"
  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = ">= 3.16.0"
    }
  }
}