service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "images" {
  path = "github.com/cloudflare/cloudflare-go/.Image"

  multiplex "CFAccount" {
    path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.AccountMultiplex"
  }

  deleteFilter "DeleteAccountFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.DeleteAccountFilter"
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
      path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ResolveAccountId"
    }
  }

  column "id" {
    description = "Image unique identifier"
  }

  column "filename" {
    description = "Image file name"
  }

  column "metadata" {
    description = "User modifiable key-value store. Can be used for keeping references to another system of record for managing images. Metadata must not exceed 1024 bytes."
    type        = "JSON"
  }

  column "requireSignedURLs" {
    description = "Indicates whether the image can be a accessed only using it's UID. If set to true, a signed token needs to be generated with a signing key to view the image."
  }

  column "variants" {
    description = "Object specifying available variants for an image."
    type        = "JSON"
  }

  column "uploaded" {
    description = "When the media item was uploaded."
  }

}