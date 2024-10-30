The CloudQuery Square plugin pulls data from [Square](https://www.squareup.com/) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://hub.cloudquery.io/plugins/destination)).

See [tables](/docs/plugins/sources/square/tables) for a list of resources supported.

## Example Configuration

This example syncs from Square to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

## Authentication

:authentication

## Configuration Reference

This is the (nested) spec used by the Square source plugin:

- `access_token` (`string`) (required):

  Your access token from Square.

- `environment` (`string`) (required):

  The environment to use. Can be `production` or `sandbox`.

- `concurrency` (`integer`) (optional) (default: `100`):

  Maximum number of requests to perform concurrently.

- `queue_size` (`integer`) (optional) (default: `10000`):

  Maximum number of items to have in the queue before waiting for an unfinished request to complete.

## Running in a Containerized Environment

To run the Square integration in a containerized environment, see the [Using CloudQuery Docker Registry Integrations Inside a Containerized Environment](https://docs.cloudquery.io/docs/advanced-topics/using-cloud-query-docker-registry-integrations-inside-a-containerized-environment) guide.