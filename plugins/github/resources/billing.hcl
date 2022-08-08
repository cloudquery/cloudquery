service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "action_billing" {
  path = "github.com/google/go-github/v45/github.ActionBilling"

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

  options {
    primary_keys = ["org"]
  }
}

resource "github" "" "package_billing" {
  path = "github.com/google/go-github/v45/github.PackageBilling"

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

  options {
    primary_keys = ["org"]
  }
}

resource "github" "" "storage_billing" {
  path = "github.com/google/go-github/v45/github.StorageBilling"

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

  options {
    primary_keys = ["org"]
  }
}