resource "digitalocean_record" "do_domain_record_a" {
  domain = digitalocean_domain.do_domain.name
  type   = "A"
  name   = "www"
  value  = "192.168.0.11"
}

resource "digitalocean_record" "do_domain_record_txt" {
  domain = digitalocean_domain.do_domain.name
  type   = "TXT"
  name   = "test"
  value  = "do_record_txt${random_id.test_id.hex}"
}

resource "digitalocean_record" "do_domain_record_ns" {
  domain = digitalocean_domain.do_domain.name
  type   = "NS"
  name   = "@"
  value  = "0.0.0.0."
}

resource "digitalocean_domain" "do_domain" {
  name = "cqdomain${random_id.test_id.hex}.com"
}