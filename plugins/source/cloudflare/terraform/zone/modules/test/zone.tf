resource "cloudflare_zone" "example" {
  zone = "cloudquery.io" // Cloudflare returns error 1049 if the account is not register, therefore we use cloudquery.io
  plan = "free"
}