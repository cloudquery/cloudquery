# CloudQuery Recipes

This section provides set of basic examples for variety of cloudquery sources and destinations plugins.

Note: Source and destinations plugins recipes are separate in this directory and you will need to combine them to run sync.

* Combine them easily by either using two different files and running `cloudquery sync aws.yml gcp.yml`

```yaml file=aws.yml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v5.0.0" # latest version of aws plugin
  tables: ["*"]
  destinations: ["postgresql"]
```

```yaml file=pg.yml
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.7.9" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

* Or combine source and destination to a single file

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v5.0.0" # latest version of aws plugin
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.7.9" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```