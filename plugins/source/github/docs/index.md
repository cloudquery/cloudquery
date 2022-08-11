# Github Provider

The CloudQuery Github provider extracts and transforms your AWS cloud assets configuration into PostgreSQL.

## Install

```shell
cloudquery init github
```

## Authentication

CloudQuery requires only a Personal Access Token, follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) by github
on how to create a personal access token for CloudQuery.

## Configuration

CloudQuery needs to be authenticated with your Github account's Personal Token in order to fetch information about your Github organizations.
CloudQuery requires only *read* permissions (we will never make any changes to your github account or organizations),
so, following the principle of least privilege, it's recommended to grant it read-only permissions.

Add the following Block your your providers list in your `cloudquery.yml` configuration.
```yaml
  - name: github
    configuration:
      access_token: "<YOUR ACCESS TOKEN HERE>
      orgs: ["<YOUR ORG NAME>"]
    resources:
      - organizations
      - repositories
      - teams
      - billing.actions
      - billing.packages
      - billing.storage
      - issues
      - hooks
      - installations
      - external_groups



```

More information can be found in the [CloudQuery documentation](https://docs.cloudquery.io/docs/developers/debugging)