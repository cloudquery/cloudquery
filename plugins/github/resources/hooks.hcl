service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "hooks" {
  path = "github.com/google/go-github/v45/github.Hook"
  options {
    primary_keys = ["id"]
  }
  ignore_columns_in_tests = ["last_response", "config"]

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

  user_relation "github" "" "deliveries" {
    path = "github.com/google/go-github/v45/github.HookDelivery"


    column "request_raw_payload" {
      generate_resolver = true
    }
    column "response_raw_payload" {
      generate_resolver = true
    }
  }

}

