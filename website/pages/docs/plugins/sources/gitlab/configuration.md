# GitLab Source Plugin Configuration Reference

## Example

This example syncs from GitLab to a Postgres destination, using API Key authentication. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: gitlab
  path: cloudquery/gitlab
  version: "VERSION_SOURCE_GITLAB"
  tables: ["*"]
  destinations: ["postgresql"]
  
  # Gitlab specific configuration
  spec:
    access_token: "<YOUR_ACCESS_TOKEN_HERE>"
    base_url: "<INSTANCE_URL>"

```

## GitLab Spec

This is the (nested) spec used by the GitLab source plugin:

- `access_token` (string, required):
  An access token for your GitLab server. Instructions on how to generate an access token [here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#create-a-personal-access-token).

- `base_url` (string, required):
  URL for your GitLab server.