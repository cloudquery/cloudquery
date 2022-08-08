service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "accounts" {
  path = "github.com/cloudflare/cloudflare-go/.Account"

  options {
    primary_keys = [
      "id"
    ]
  }

  column "id" {
    description = "The unique universal identifier for a Cloudflare account."
  }

  column "name" {
    description = "Cloudflare account name."
  }

  column "type" {
    description = "Cloudflare account type."
  }

  column "created_on" {
    description = "Creation timestamp of the account."
  }

  column "enforce_two_factor" {
    description = "True if the account has enforce 2fa authentication."
  }

  column "settings" {
    skip_prefix = true
  }

  user_relation "cloudflare" "" "account_members" {
    path = "github.com/cloudflare/cloudflare-go/.AccountMember"

    userDefinedColumn "account_id" {
      description = "The Account ID of the resource."
      type        = "string"
      resolver "resolveCFAccount" {
        path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
      }
    }

    column "id" {
      description = "The unique universal identifier for a Cloudflare account member."
    }

    column "user_id" {
      description = "Cloudflare user id."
    }

    column "user_first_name" {
      description = "Cloudflare user first name."
    }

    column "user_last_name" {
      description = "Cloudflare user last name."
    }

    column "user_email" {
      description = "Cloudflare user email."
    }

    column "user_two_factor_authentication_enabled" {
      description = "True if user has enabled 2fa authentication."
    }

    column "status" {
      description = "Cloudflare account member status."
    }

    relation "cloudflare" "" "roles" {
      path = "github.com/cloudflare/cloudflare-go/.AccountRole"

      userDefinedColumn "account_id" {
        description = "The Account ID of the resource."
        type        = "string"
        resolver "resolveCFAccount" {
          path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
        }
      }

      column "id" {
        description = "The unique universal identifier for a Cloudflare account member role."
      }

      column "name" {
        description = "Cloudflare account member role name."
      }

      column "description" {
        description = "Cloudflare account member role description."
      }

      column "permissions" {
        description = "Cloudflare account member role permissions."
      }
    }
  }
}
