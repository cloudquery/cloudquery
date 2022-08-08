service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "external_groups" {
  path = "github.com/google/go-github/v45/github.ExternalGroup"
  options {
    primary_keys = ["org", "group_id"]
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
}

