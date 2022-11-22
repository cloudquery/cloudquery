# CloudQuery Recipes

This section provides set of basic examples for variety of cloudquery sources and destinations plugins.

Note: Source and destinations plugins recipes are separate in this directory and you will need to combine them to run sync.

* Combine them easily by either using two different files and running `cloudquery sync aws.yml gcp.yml`

```yaml copy
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v7.1.4" # latest version of aws plugin
  tables: ["*"]
  destinations: ["postgresql"]
```

```yaml copy
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.7.13" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

* Or combine source and destination to a single file

```yaml copy
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v7.1.4" # latest version of aws plugin
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.7.13" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```