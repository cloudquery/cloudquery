# GitHub Plugin

The CloudQuery GitHub plugin extracts your GitHub API.

## Authentication

CloudQuery requires only a Personal Access Token. follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to create a personal access token for CloudQuery.

CloudQuery requires only *read* permissions (we will never make any changes to your GitHub account or organizations),
so, following the principle of least privilege, it's recommended to grant it read-only permissions.

## Configuration

To configure CloudQuery to extract from github, create a `.yml` file in your CloudQuery configuration directory.
For example, the following configuration will extract information from the `cloudquery` organization, and connect it to a `postgresql` destination plugin

```yaml
kind: source
spec:
  # Source spec section
  name: github
  path: cloudquery/github
  version: "VERSION_SOURCE_GCP"
  tables: ["*"]
  destinations: ["postgresql"]
  spec:
    access_token: <YOUR_ACCESS_TOKEN_HERE> # Required. Personal Access Token
    orgs: ["cloudquery"] # Required. List of organizations to extract from
```
