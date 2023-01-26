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

- `max_requests_per_second` (int) (default: 10)
  PagerDuty API is heavily rate-limited (900 requests/min = 15 requests/sec, across the entire organization). 
  This option allows you to control the rate at which the plugin will make requests to the API. 
  You can reduce this parameter in case you are still seeing rate limit errors (status code 429), or increase
  it if your PagerDuty API quota is higher. See https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUz-rate-limiting#what-are-our-limits for more info.