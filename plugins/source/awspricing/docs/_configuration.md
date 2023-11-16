```yaml copy
kind: source
spec:
  name: "awspricing"
  path: "cloudquery/awspricing"
  registry: "cloudquery"
  version: "VERSION_SOURCE_AWSPRICING"
  tables: ["*"]
  destinations:
    - "DESTINATION_NAME"
  spec:
    # Optional parameters
    # region_codes: []
    # offer_codes: []
    # concurrency: 10000
```
