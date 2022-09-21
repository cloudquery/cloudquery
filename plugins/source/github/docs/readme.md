# GitHub Plugin

The CloudQuery GitHub plugin extracts your GitHub information, normalizes them and stores them in into any of CloudQuery destination plugins.

## Install

```bash
cloudquery gen source github
```

## Authentication

CloudQuery requires GitHub Personal Access Token. follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) by GitHub
on how to create a personal access token for CloudQuery.

## Configuration

CloudQuery needs to be authenticated with your GitHub account's Personal Token in order to fetch information about your GitHub organizations.
CloudQuery requires only _read_ permissions (we will never make any changes to your GitHub account or organizations), so, following the principle of least privilege, it's recommended to grant it read-only permissions.

Add the following block to your providers list in your `github.yml` configuration. CloudQuery will fetch information about all the organizations specified in `orgs`.

```yaml title="github.yml"
kind: "source"
spec:
  name: "github"
  tables: ["*"]
  spec:
    access_token: "<YOUR ACCESS TOKEN HERE>"
    orgs: ["<YOUR ORG NAME>"]
```

After that, edit `github.yml` and set the `access_token` and `orgs` values.
CloudQuery will fetch information about all the organizations specified in `orgs`.

More information can be found in the [CloudQuery documentation](https://docs.cloudquery.io/docs/intro)
