# GCP Source Plugin Recipes

Full spec options for GCP Source available [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/docs/configuration.md).

```yaml copy
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "v3.0.2" # latest version of gcp plugin
  tables: ["*"]
  destinations: ["YOUR_DESTINATION"]
```
