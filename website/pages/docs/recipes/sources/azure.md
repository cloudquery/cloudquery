# Azure Source Plugin Recipes

Full spec options for the Azure source plugin are available [here](/docs/plugins/sources/azure/configuration#azure-spec).

```yaml copy
kind: source
spec:
  name: azure
  path: cloudquery/azure
  version: "VERSION_SOURCE_AZURE"
  tables: ["*"]
  destinations: ["<destination>"]
```
