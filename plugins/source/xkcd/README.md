# CloudQuery XKCD Source Plugin

This CloudQuery source plugin fetches data from the [XKCD API](https://xkcd.com/json.html), allowing you to load the XKCD comic data into any CloudQuery-supported destination (e.g. PostgreSQL, Elasticsearch, Snowflake, etc.). See [CloudQuery destinations](https://www.cloudquery.io/docs/plugins/destinations/overview) for a complete list of supported destinations.

It was originally developed as part of a live-coding tutorial on how to write your own CloudQuery source plugin. It only took 30 minutes! You can watch the video here: https://www.youtube.com/watch?v=3Ka_Ob8E6P8

## Links

 - [Video Tutorial](https://www.youtube.com/watch?v=3Ka_Ob8E6P8)
 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)

## Configuration

The following source configuration file will sync all comics to a local SQLite database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: xkcd
  path: hermanschaaf/xkcd
  version: v2.0.0
  destinations: ["sqlite"]
  tables: ["*"]
  spec:
---
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  version: v2.2.0
  spec:
    connection_string: ./db.sqlite
```

Save the file as `config.yaml`, then run:

```
cloudquery sync config.yaml
```
