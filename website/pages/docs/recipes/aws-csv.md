# AWS + CSV

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v4.1.0" # latest version of aws plugin
  tables: ["*"]
  destinations: ["csv"]
---
kind: destination
spec:
  name: csv
  path: cloudquery/csv
  write_mode: "append"
  version: "v1.0.0" # latest version of csv plugin
```
