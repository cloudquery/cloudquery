```yaml copy
kind: source
spec:
  name: "jira"
  path: "cloudquery/jira" # Jira is a premium plugin: https://cloudquery.io/buy/jira
  registry: "cloudquery"
  version: "VERSION_SOURCE_JIRA"
  destinations: ["DESTINATION_NAME"]
  tables: ["*"]
  spec:
    base_url: "https://your_uri.atlassian.net"
    username: "your_username"
    token: ${ENV_WITH_API_KEY}
```
