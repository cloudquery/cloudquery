service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "organizations" {
  path = "github.com/google/go-github/v45/github.Organization"
  options {
    primary_keys = ["id"]
  }
  multiplex "OrgMultiplex" {
    path = "github.com/cloudquery/cq-provider-github/client.OrgMultiplex"
  }


  column "text_matches" {
    type              = "json"
    generate_resolver = true
  }

  user_relation "github" "" "members" {
    path = "github.com/google/go-github/v45/github.User"

    options {
      primary_keys = ["org", "id"]
    }

    userDefinedColumn "org" {
      type        = "string"
      description = "The Github Organization of the resource."
      resolver "resolveOrg" {
        path = "github.com/cloudquery/cq-provider-github/client.ResolveOrg"
      }
    }


    column "text_matches" {
      type              = "json"
      generate_resolver = true
    }

    user_relation "github" "" "membership" {
      path = "github.com/google/go-github/v45/github.Membership"
      column "user" {
        skip = true
      }
      column "organization" {
        skip = true
      }
    }
  }

}

