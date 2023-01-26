# CloudQuery PagerDuty Source Plugin Configuration Reference

## Example

In order to get started with the PagerDuty plugin, you need to create a YAML file in your CloudQuery configuration directory (e.g. named `pagerduty.yml`).

This example connects a single PagerDuty subscription to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
spec:
  # Source spec section
  name: "pagerduty"
  path: "cloudquery/pagerduty"
  version: "VERSION_SOURCE_PAGERDUTY"
  destinations: ["postgresql"]

  spec:
    # optional
    team_ids: ["<team_id>"]
```

## PagerDuty Spec

This is the (nested) spec used by the PagerDuty source plugin.

- `team_ids` ([]string) (default: empty. will sync data from all available teams)

  If specified, limits the sync to only resources related to the specified teams.