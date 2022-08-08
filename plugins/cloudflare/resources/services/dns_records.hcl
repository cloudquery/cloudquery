service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "dns_records" {
  path = "github.com/cloudflare/cloudflare-go/.DNSRecord"

  options {
    primary_keys = [
      "id"
    ]
  }

  userDefinedColumn "account_id" {
    description = "The Account ID of the resource."
    type        = "string"
    resolver "resolveCFAccount" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
    }
  }

  column "created_on" {
    description = "When the record was created."
  }

  column "modified_on" {
    description = "When the record was last modified."
  }

  column "type" {
    description = "Record type."
  }

  column "name" {
    description = "DNS record name."
  }

  column "content" {
    description = "A valid IPv4 address."
  }

  column "id" {
    description = "DNS record identifier tag."
  }

  column "zone_id" {
    description = "Zone identifier tag."
  }

  column "zone_name" {
    description = "The domain of the record."
  }

  column "priority" {
    description = "The priority of the record."
  }

  column "ttl" {
    description = "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'"
  }

  column "proxied" {
    description = "Whether the record is receiving the performance and security benefits of Cloudflare."
  }

  column "proxiable" {
    description = "Whether the record can be proxied by Cloudflare or not."
  }

  column "locked" {
    description = "Whether this record can be modified/deleted (true means it's managed by Cloudflare)."
  }

  column "meta" {
    description = "Extra Cloudflare-specific information about the record."
    type        = "JSON"
  }

  column "data" {
    description = "Metadata about the record."
    type        = "JSON"
  }

}
