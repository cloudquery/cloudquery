# CloudQuery Azure Source Plugin Configuration Reference

## Example

This example connects a single Azure subscription to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
spec:
  # Source spec section
  name: "azure"
  path: "cloudquery/azure"
  version: "VERSION_SOURCE_AZURE"
  destinations: ["postgresql"]

  spec:
    # Azure Spec section described below
    subscriptions: ["00000000-0000-0000-0000-000000000000"]
```

## Azure Spec

This is the (nested) spec used by the Azure source plugin.

- `subscriptions` ([]string) (default: empty. Will use all visible subscriptions)

  Specify which subscriptions to sync data from.
