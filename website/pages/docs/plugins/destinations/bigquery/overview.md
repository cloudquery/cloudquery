# CloudQuery BigQuery Destination Plugin

The BigQuery plugin helps you sync data to a BigQuery database running on GCP.

The plugin currently only supports a streaming mode, which is suitable for small- to medium-sized datasets. It will stream the results directly to the BigQuery database. There is no additional setup needed apart from authentication to BigQuery.

A batch mode of operation is also being developed to support larger datasets, but this is not currently supported.

## Configuration



## BigQuery Spec

This is the top level spec used by the BigQuery destination plugin.

- `project_id` (string) (required)
  The id of the project where the destination BigQuery database resides.

- `dataset_id` (string) (required)
  The id of the BigQuery dataset within the project. This dataset needs to be created before running a sync or migration.

## Underlying library

We use the official [cloud.google.com/go/bigquery](https://pkg.go.dev/cloud.google.com/go/bigquery) package for database connection.
