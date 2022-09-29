# GitHub Plugin

The CloudQuery GitHub plugin extracts your GitHub information.

## Install


## Authentication

CloudQuery requires only a Personal Access Token. follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) by GitHub
on how to create a personal access token for CloudQuery.

## Configuration

CloudQuery needs to be authenticated with your GitHub account's Personal Token in order to fetch information about your GitHub organizations.
CloudQuery requires only *read* permissions (we will never make any changes to your GitHub account or organizations),
so, following the principle of least privilege, it's recommended to grant it read-only permissions.

First, generate the configuration template using the following command:
```bash
cloudquery gen source github
```

After that, edit `github.yml` and set the `access_token` and `orgs` values. 
CloudQuery will fetch information about all the organizations specified in `orgs`.

More information can be found in the [CloudQuery documentation](https://docs.cloudquery.io/docs/intro)
