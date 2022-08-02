service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "repositories" {
  path = "github.com/google/go-github/v45/github.Repository"

  multiplex "OrgMultiplex" {
    path = "github.com/cloudquery/cq-provider-github/client.OrgMultiplex"
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

  column "s_v_n_url" {
    type = "string"
    rename = "svn_url"
  }

  column "use_squash_p_r_title_as_default" {
    type = "bool"
    rename = "use_squash_pr_title_as_default"
  }

  column "license_s_p_d_x_id" {
    type = "string"
    rename = "license_spdx_id"
  }

  column "text_matches" {
    type              = "json"
    generate_resolver = true
  }

  column "owner_text_matches" {
    type              = "json"
    generate_resolver = true
  }

  options {
    primary_keys = ["id"]
  }

}

