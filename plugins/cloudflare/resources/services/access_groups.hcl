service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "access_groups" {
  path = "github.com/cloudflare/cloudflare-go/.AccessGroup"

  multiplex "CFZone" {
    path   = "github.com/cloudquery/cq-provider-cloudflare/client.ZoneMultiplex"
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
    description = "Zone identifier tag."
    type        = "string"
    resolver "resolveCFZone" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveZoneId"
    }
  }

  column "id" {
    description = "The unique identifier for the Access group."
  }

  column "created_at" {
    description = "Hashed script content, can be used in a If-None-Match header when updating."
  }

  column "updated_at" {
    description = "Size of the script, in bytes."
  }

  column "name" {
    description = "The name of the Access group."
  }

  column "include" {
    description = "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules."
    type        = "JSON"
  }

  column "exclude" {
    description = "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules."
    type        = "JSON"
  }

  column "require" {
    description = "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules."
    type        = "JSON"
  }

}