```yaml copy
kind: source
spec:
  # Source spec section
  name: "pagerduty"
  path: "cloudquery/pagerduty"
  version: "VERSION_SOURCE_PAGERDUTY"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec:
    # optional
    team_ids: ["<team_id>"]
