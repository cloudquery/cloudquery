# SQLite Destination Plugin Recipes

Full spec options for CSV destination available [here](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/csv).


## Basic

This is a basic configuration that will save all your sync resources to `db.sql`.

```yaml
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  version: "v1.0.5" # latest version of sqlite plugin
  spec:
    connection_string: ./db.sql
```

Now you can easily explore the data locally with SQLite CLI without running any database.
