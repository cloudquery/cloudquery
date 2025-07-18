This CloudQuery source plugin fetches data from the [XKCD API](https://xkcd.com/json.html), allowing you to load the XKCD comic data into any CloudQuery-supported destination (e.g. PostgreSQL, Elasticsearch, Snowflake, etc.). See [CloudQuery destinations](https://hub.cloudquery.io/plugins/destination) for a complete list of supported destinations.

It was originally developed as part of a live-coding tutorial on how to write your own CloudQuery source plugin. It only took 30 minutes! You can watch the video here: https://www.youtube.com/watch?v=3Ka_Ob8E6P8

## Links

 - [Video Tutorial](https://www.youtube.com/watch?v=3Ka_Ob8E6P8)
 - [CloudQuery Quickstart Guide](https://cli-docs.cloudquery.io/docs/quickstart)

## Configuration

The following configuration syncs from the XKCD API to a Postgres destination.

The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec). The config for the `postgresql` destination is not shown here. See our [Quickstart](https://cli-docs.cloudquery.io/docs/quickstart) if you need help setting up the destination.

:configuration
