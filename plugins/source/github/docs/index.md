# GitHub Plugin

The CloudQuery GitHub provider extracts and transforms your GitHub information into PostgreSQL.

## Install

```bash
cloudquery init github
```

## Authentication

CloudQuery requires only a Personal Access Token. follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) by GitHub
on how to create a personal access token for CloudQuery.

## Configuration

CloudQuery needs to be authenticated with your GitHub account's Personal Token in order to fetch information about your GitHub organizations.
CloudQuery requires only *read* permissions (we will never make any changes to your GitHub account or organizations),
so, following the principle of least privilege, it's recommended to grant it read-only permissions.

Add the following block to your providers list in your `cloudquery.yml` configuration. CloudQuery will
fetch information about all the organizations specified in `orgs`.

```yaml
- name: github
  configuration:
    access_token: "<YOUR ACCESS TOKEN HERE>"
    orgs: ["<YOUR ORG NAME>"]
  resources:
    - "*"
```

More information can be found in the [CloudQuery documentation](https://docs.cloudquery.io/docs/intro)
