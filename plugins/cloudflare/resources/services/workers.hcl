service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "workers_scripts" {
  path = "github.com/cloudflare/cloudflare-go/.WorkerMetaData"

  multiplex "CFAccount" {
    path   = "github.com/cloudquery/cq-provider-cloudflare/client.AccountMultiplex"
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

  column "id" {
    description = "The unique universal identifier for a Cloudflare zone."
  }

  column "etag" {
    rename = "etag"
    description = "Hashed script content, can be used in a If-None-Match header when updating."
  }

  column "size" {
    description = "Size of the script, in bytes."
  }

  column "created_on" {
    description = "When the script was created."
  }

  column "modified_on" {
    description = "When the script was last modified."
  }

  user_relation "cloudflare" "" "cron_triggers" {
    path = "github.com/cloudflare/cloudflare-go/.WorkerCronTrigger"

    userDefinedColumn "account_id" {
      description = "The Account ID of the resource."
      type        = "string"
      resolver "resolveCFAccount" {
        path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
      }
    }

    column "cron" {
      description = "Raw cron expression"
    }

    column "created_on" {
      description = "When the Cron was created"
    }

    column "modified_on" {
      description = "When the Cron was last modified"
    }
  }

  user_relation "cloudflare" "" "secrets" {
    path = "github.com/cloudflare/cloudflare-go/.WorkersSecret"

    userDefinedColumn "account_id" {
      description = "The Account ID of the resource."
      type        = "string"
      resolver "resolveCFAccount" {
        path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
      }
    }

    column "name" {
      description = "Secret name"
    }

    column "type" {
      description = "Secret type"
    }
  }

#
#  column "script" {
#    description = "Raw script content, as a string."
#  }
#
#  column "usage_model" {
#    description = "Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound')."
#  }
}


resource "cloudflare" "" "workers_routes" {
  path = "github.com/cloudflare/cloudflare-go/.WorkerRoute"

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
    description = "The Zone ID of the resource."
    type        = "string"
    resolver "resolveCFZone" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveZoneId"
    }
  }


  column "id" {
    description = "API item identifier tag"
  }

  column "pattern" {
    description = "The pattern of the route."
  }

  column "enabled" {
    description = "Whether the route is enabled"
  }

  column "script" {
    description = "Name of the script to apply when the route is matched. The route is skipped when this is blank/missing."
  }
}