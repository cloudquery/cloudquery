```yaml copy
kind: source
spec:
  # Source spec section
  name: k8s
  path: cloudquery/k8s
  version: "VERSION_SOURCE_K8S"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec:
    contexts: ["context"]
```
