```yaml copy
kind: source
spec:
  # Source spec section
  name: k8s
  path: cloudquery/k8s
  registry: cloudquery
  version: "VERSION_SOURCE_K8S"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # Learn more about the configuration options at https://cql.ink/k8s_source
  spec:
    contexts: ["context"]
```
