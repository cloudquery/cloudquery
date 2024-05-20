```yaml copy
kind: source
spec:
  name: "xkcd"
  path: "cloudquery/xkcd"
  version: "VERSION_SOURCE_XKCD"
  tables: ["xkcd_comics"]
  destinations:
    - "DESTINATION_NAME"
  spec:
```
