# GCP Source Plugin Recipes

Full spec options for GCP Source available [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/docs/configuration.md).

```yaml copy
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "${VERSION_DESTINATION_GCP}"
  tables: ["*"]
  destinations: ["YOUR_DESTINATION"]
```
