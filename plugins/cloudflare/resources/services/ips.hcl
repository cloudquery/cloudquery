service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "ips" {
  userDefinedColumn "ip" {
    type = "string"
    description = "Cloudflare ip cidr address."
  }

  userDefinedColumn "type" {
    type = "string"
    description = "Ip type, ipv4, ipv6, ipv4_china, ipv6_china."
  }

}
