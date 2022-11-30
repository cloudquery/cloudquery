import { Callout } from 'nextra-theme-docs'
import { getLatestVersion } from "../../../../../utils/versions";

# BigQuery Destination Plugin

The BigQuery plugin helps you sync data to a BigQuery database running on GCP.

The plugin currently only supports a streaming mode through the legacy streaming API. This is suitable for small- to medium-sized datasets, and will stream the results directly to the BigQuery database. A batch mode of operation is also being developed to support larger datasets, but this is not currently supported.

<Callout type="info">
Streaming is not available for the [Google Cloud free tier](https://cloud.google.com/bigquery/pricing#free-tier).
</Callout>

## Before you begin

1. Make sure that billing is enabled for your Cloud project. Learn how to [check if billing is enabled on a project](https://cloud.google.com/billing/docs/how-to/verify-billing-enabled).
2. Create a BigQuery dataset that will contain the tables synced by CloudQuery. CloudQuery will automatically create the tables as part of a migration run on the first `sync`.
3. Ensure that you have write access to the dataset. See [Required Permissions](https://cloud.google.com/bigquery/docs/streaming-data-into-bigquery) for details.

## Configuration

See an example configuration for the BigQuery destination under [recipes](/docs/recipes/destinations/bigquery).

The BigQuery plugin supports [all three write modes](/docs/reference/destination-spec#write_mode): `append`, `overwrite` and `overwrite-delete-stale`.

<Callout type="info">
When using the `overwrite-delete-stale` write mode, syncs must be spaced at least 90 minutes apart. Otherwise, the delete-stale functionality will not be able to delete rows from the more recent syncs. This is because BigQuery's internal streaming buffer disallows deletion of rows in the buffer. For more information, see [BigQuery Limitations](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language).
</Callout>


## BigQuery Spec

This is the top level spec used by the BigQuery destination plugin.

- `project_id` (string) (required)

  The id of the project where the destination BigQuery database resides.


- `dataset_id` (string) (required)

  The id of the BigQuery dataset within the project. This dataset needs to be created before running a sync or migration.

## Underlying library

We use the official [cloud.google.com/go/bigquery](https://pkg.go.dev/cloud.google.com/go/bigquery) package for database connection.
