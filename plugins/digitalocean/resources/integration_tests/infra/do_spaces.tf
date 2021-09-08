resource "digitalocean_spaces_bucket" "do_spaces" {
  name   = "do-spaces-${var.test_prefix}-${var.test_suffix}"
  region = "nyc3"
  acl    = "public-read"


  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
    max_age_seconds = 3000
  }

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["PUT", "POST", "DELETE"]
    allowed_origins = ["https://www.${var.test_prefix}-${var.test_suffix}.com"]
    max_age_seconds = 3000
  }


}
