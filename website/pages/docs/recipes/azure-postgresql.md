# Azure + PostgreSQL

```yaml
kind: source
spec:
  name: azure
  path: cloudquery/azure
  version: v1.0.1 # latest version of azure plugin
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: v1.1.0 # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```