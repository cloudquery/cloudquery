service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "issues" {
  path = "github.com/google/go-github/v45/github.Issue"
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

  column "assignee_text_matches" {
    type = "json"
    generate_resolver = true
  }

  column "milestone_creator_text_matches" {
    type = "json"
    generate_resolver = true
  }

  column "user_text_matches" {
    type = "json"
    generate_resolver = true
  }

  column "closed_by_text_matches" {
    type = "json"
    generate_resolver = true
  }

  column "repository" {
    type = "int"
    rename = "repository_id"
    generate_resolver = true
  }

  column "text_matches" {
    type              = "json"
    generate_resolver = true
  }


  relation "github" "" "assignees" {
    path = "github.com/google/go-github/v45/github.User"

    options {
      primary_keys = ["issue_id", "id"]
    }

    userDefinedColumn "issue_id" {
      type        = "int"
      description = "The id of the issue"
      resolver "parentPathResolver" {
        path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        path_resolver = true
        params = ["id"]
      }
    }

    column "text_matches" {
      type              = "json"
      generate_resolver = true
    }

  }

}

