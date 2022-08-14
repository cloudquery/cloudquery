resource "awslightsail_certificate" "awslightsail_certificate" {
  name                      = "${var.prefix}_awslightsail_certificate"
  domain_name               = "example.com"
  subject_alternative_names = ["www.example.com"]
}