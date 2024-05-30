The CloudQuery Bitbucket plugin pulls data from [Bitbucket](https://bitbucket.org/) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://hub.cloudquery.io/plugins/destination)).

The plugin discover all workspaces and repositories in your account and syncs them to the destination.

## Example Configuration

This example syncs from Bitbucket to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

## Authentication

:authentication

## Example

```yaml
kind: source
spec:
  name: bitbucket
  registry: docker
  path: ghcr.io/cloudquery/cq-source-bitbucket:VERSION_SOURCE_BITBUCKET
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # bitbucket-specific configuration
  spec:
    username: "${BITBUCKET_USERNAME}" # required
    password: "${BITBUCKET_PASSWORD}" # required
  ...
```

## Configuration Reference

This is the (nested) spec used by the Bitbucket source plugin:

- `username` (`string`) (required):

  The Bitbucket username associated with the [app password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/).

- `password` (`string`) (required):

  The Bitbucket password associated with the [app password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/).

## Running in a Containerized Environment

To run the Bitbucket plugin in a containerized environment, see the [Using CloudQuery Docker Registry Plugins Inside a Containerized Environment](https://docs.cloudquery.io/docs/advanced-topics/using-cloud-query-docker-registry-plugins-inside-a-containerized-environment) guide.