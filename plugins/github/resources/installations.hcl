service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "installations" {
  path = "github.com/google/go-github/v45/github.Installation"
  options {
    primary_keys = ["id"]
  }
  multiplex "OrgMultiplex" {
    path = "github.com/cloudquery/cq-provider-github/client.OrgMultiplex"
  }

  userDefinedColumn "org" {
    type        = "string"
    description = "The Github Organization of the resource."
    resolver "resolveOrg" {
      path = "github.com/cloudquery/cq-provider-github/client.ResolveOrg"
    }
  }
  column "account_text_matches" {
    type = "json"
    generate_resolver = true
  }
  column "suspended_by_text_matches" {
    type = "json"
    generate_resolver = true
  }
}

