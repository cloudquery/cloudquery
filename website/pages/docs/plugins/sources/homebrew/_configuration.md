```yaml copy
kind: source
spec:
  name: "homebrew"
  path: "cloudquery/homebrew"
  registry: "cloudquery"
  version: "VERSION_SOURCE_HOMEBREW"
  tables: ["*"]
  destinations:
    - "DESTINATION_NAME"
  backend_options:
    table_name: "cq_state_homebrew"
    connection: "@@plugins.DESTINATION_NAME.connection"
  spec:
```
