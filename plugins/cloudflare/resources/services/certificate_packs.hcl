service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "certificate_packs" {
  path = "github.com/cloudflare/cloudflare-go/.CertificatePack"

  multiplex "CFZone" {
    path = "github.com/cloudquery/cq-provider-cloudflare/client.ZoneMultiplex"
  }

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

  userDefinedColumn "zone_id" {
    description = "The Zone ID of the resource."
    type        = "string"
    resolver "resolveCFZone" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveZoneId"
    }
  }

  column "id" {
    description = "The unique identifier for a certificate_pack"
  }

  column "type" {
    description = "Type of certificate pack"
  }

  column "hosts" {
    description = "comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty."
  }

  column "primary_certificate" {
    description = "Identifier of the primary certificate in a pack"
  }

  column "certificates" {
    description = "list of the certificates in the pack"
  }
}