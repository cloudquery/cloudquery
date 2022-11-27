# Azure Source Plugin Recipes

Full spec options for Azure Source available [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/azure/docs/configuration.md).

```yaml copy
kind: source
spec:
  name: azure
  path: cloudquery/azure
  version: "${VERSION_SOURCE_AZURE}"
  tables: ["*"]
  destinations: ["<destination>"]
```
