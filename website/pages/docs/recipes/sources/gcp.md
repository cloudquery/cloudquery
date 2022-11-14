# GCP Source Plugin Recipes

Full spec options for GCP Source available [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/docs/configuration.md).

```yaml
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "v2.4.16" # latest version of gcp plugin
  tables: ["*"]
  destinations: ["YOUR_DESTINATION"]
```
