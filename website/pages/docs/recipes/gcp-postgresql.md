# GCP + PostgreSQL

```yaml
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "1.0.3" # latest version of gcp plugin
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.3.5" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```
