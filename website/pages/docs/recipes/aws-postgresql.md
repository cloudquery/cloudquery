# AWS + PostgreSQL

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v4.1.0" # latest version of aws plugin
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.6.0" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```
