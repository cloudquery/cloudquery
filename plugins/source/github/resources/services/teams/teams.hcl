//check-for-changes
service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "teams" {
  path = "github.com/google/go-github/v45/github.Team"

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

  column "text_matches" {
    type              = "json"
    generate_resolver = true
  }

  column "parent" {
    type              = "int"
    generate_resolver = true
  }

  column "organization" {
    skip = true
  }

  column "l_d_a_p_d_n" {
    rename = "ldapdn"
  }

  options {
    primary_keys = ["org", "id"]
  }

  user_relation "github" "" "members" {
    path = "github.com/google/go-github/v45/github.User"

    options {
      primary_keys = ["team_id", "id"]
    }

    userDefinedColumn "team_id" {
<<<<<<<< HEAD:plugins/source/github/resources/teams.hcl
      type = "string"
========
      type        = "int"
>>>>>>>> 7c2a2f51a (chore(build): Add support for drift detection of generated code (#22)):resources/services/teams/teams.hcl
      //argument ("name")
      description = "The id of the team"
      resolver "parentPathResolver" {
        path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params        = ["id"]
      }
    }
    userDefinedColumn "org" {
      type        = "string"
      description = "The Github Organization of the resource."
      resolver "resolveOrg" {
        path = "github.com/cloudquery/cloudquery/plugins/source/github/client.ResolveOrg"
      }
    }


    column "text_matches" {
      type              = "json"
      generate_resolver = true
    }

  }

  user_relation "github" "" "repositories" {
    path = "github.com/google/go-github/v45/github.Repository"

    options {
      primary_keys = ["team_id", "id"]
    }

    userDefinedColumn "team_id" {
<<<<<<<< HEAD:plugins/source/github/resources/teams.hcl
      type = "string"
========
      type        = "int"
>>>>>>>> 7c2a2f51a (chore(build): Add support for drift detection of generated code (#22)):resources/services/teams/teams.hcl
      //argument ("name")
      description = "The id of the team"
      resolver "parentPathResolver" {
        path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params        = ["id"]
      }
    }
    column "team_id" {
      skip = true
    }

    column "parent" {
      type              = "int"
      generate_resolver = true
    }

    column "source" {
      type              = "int"
      generate_resolver = true
    }

    column "template_repository" {
      type              = "int"
      generate_resolver = true
    }


    column "text_matches" {
      type              = "json"
      generate_resolver = true
    }

    column "owner_text_matches" {
      type              = "json"
      generate_resolver = true
    }
  }

}