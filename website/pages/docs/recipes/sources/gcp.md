# GCP Source Plugin Recipes

Full spec options for the GCP source plugin are available [here](/docs/plugins/sources/gcp/configuration#gcp-spec).

```yaml copy
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "VERSION_SOURCE_GCP"
  tables: ["*"]
  destinations: ["YOUR_DESTINATION"]
```
