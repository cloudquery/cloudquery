//check-for-changes
service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "issues" {
  path = "github.com/google/go-github/v45/github.Issue"
  options {
    primary_keys = ["org", "id"]
  }

  multiplex "OrgMultiplex" {
    path = "github.com/cloudquery/cloudquery/plugins/source/github/client.OrgMultiplex"
  }
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/cq-provider-github/client.IgnoreError"
  }

  userDefinedColumn "org" {
    type        = "string"
    description = "The Github Organization of the resource."
    resolver "resolveOrg" {
      path = "github.com/cloudquery/cloudquery/plugins/source/github/client.ResolveOrg"
    }
  }

  column "assignee_text_matches" {
    type              = "json"
    generate_resolver = true
  }

  column "milestone_creator_text_matches" {
    type              = "json"
    generate_resolver = true
  }

  column "user_text_matches" {
    type              = "json"
    generate_resolver = true
  }

  column "closed_by_text_matches" {
    type              = "json"
    generate_resolver = true
  }

  column "repository" {
    type              = "int"
    rename            = "repository_id"
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
<<<<<<<< HEAD:plugins/source/github/resources/issues.hcl
        path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        path_resolver = true
        params        = ["id"]
========
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params = ["id"]
>>>>>>>> 7c2a2f51a (chore(build): Add support for drift detection of generated code (#22)):resources/services/issues/issues.hcl
      }
    }

    column "text_matches" {
      type              = "json"
      generate_resolver = true
    }

  }

}

