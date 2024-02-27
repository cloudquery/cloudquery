```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: gitlab
  path: cloudquery/gitlab
  registry: cloudquery
  version: "VERSION_SOURCE_GITLAB"
  tables: ["gitlab_users"]
  destinations: ["DESTINATION_NAME"]

  # Gitlab specific configuration
  spec:
    # required
    access_token: "${GITLAB_ACCESS_TOKEN}"
    # optional, leave empty for GitLab SaaS
    # base_url: "<INSTANCE_URL>"
```
