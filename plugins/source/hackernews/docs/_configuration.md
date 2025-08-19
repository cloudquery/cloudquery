```yaml copy
kind: source
spec:
  name: "hackernews"
  path: "cloudquery/hackernews"
  registry: "cloudquery"
  version: "VERSION_SOURCE_HACKERNEWS"
  tables: ["*"]
  backend_options:
    table_name: "cq_state_hackernews"
    connection: "@@plugins.DESTINATION_NAME.connection"
  destinations:
    - "DESTINATION_NAME"
  # Learn more about the configuration options at https://cql.ink/hackernews_source
  spec:
    item_concurrency: 100
    start_time: 3 hours ago
```
