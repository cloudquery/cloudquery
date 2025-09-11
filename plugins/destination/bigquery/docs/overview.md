---
name: BigQuery
stage: GA
title: BigQuery Destination Plugin
description: CloudQuery BigQuery destination plugin documentation
---
# BigQuery Destination Plugin

:badge

The BigQuery plugin syncs data from any CloudQuery source plugin(s) to a BigQuery database running on Google Cloud Platform.

The plugin currently only supports a streaming mode through the legacy streaming API. This is suitable for small- to medium-sized datasets, and will stream the results directly to the BigQuery database. A batch mode of operation is being developed to support larger datasets, but this is not currently supported.

:::callout{type="info"}
Streaming is not available for the [Google Cloud free tier](https://cloud.google.com/bigquery/pricing#free-tier).
:::

## Before you begin

1. Make sure that billing is enabled for your Cloud project. Learn how to [check if billing is enabled on a project](https://cloud.google.com/billing/docs/how-to/verify-billing-enabled).
2. Create a BigQuery dataset that will contain the tables synced by CloudQuery. CloudQuery will automatically create the tables as part of a migration run on the first `sync`.
3. Ensure that you have write access to the dataset. See [Required Permissions](https://cloud.google.com/bigquery/docs/streaming-data-into-bigquery) for details.

## Example config

:configuration

The BigQuery destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

Note that the BigQuery plugin only supports the `append` write mode. 

## Authentication

:authentication

## BigQuery Spec

This is the top-level spec used by the BigQuery destination plugin.

- `project_id` (`string`) (required)

  The id of the project where the destination BigQuery database resides.


- `dataset_id` (`string`) (required)

  The name of the BigQuery dataset within the project, e.g. `my_dataset`.
  This dataset needs to be created before running a sync or migration.


- `dataset_location` (`string`) (optional)

  The data location of the BigQuery dataset. If set, will be used as the default location for job operations.
  Pro-tip: this can solve "dataset not found" issues for newly created datasets.


- `time_partitioning` (`string`) (options: `none`, `hour`, `day`) (default: `none`)

  The time partitioning to use when creating tables. The partition time column used will always be `_cq_sync_time` so that all rows for a sync run will be partitioned on the hour/day the sync started.

- `service_account_key_json` (`string`) (optional) (default: empty).

  GCP service account key content.
  This allows for using different service accounts for the GCP source and BigQuery destination.
  If using service account keys, it is best to use [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).

- `endpoint` (`string`) (optional)

  The BigQuery API endpoint to use. This is useful for testing against a local emulator.

- `batch_size` (`integer`) (optional) (default: `10000`)

  Number of records to write before starting a new object.

- `batch_size_bytes` (`integer`) (optional) (default: `5242880` (5 MiB))

  Number of bytes (as Arrow buffer size) to write before starting a new object.

- `batch_timeout` (`duration`) (optional) (default: `10s` (10 seconds))

  Maximum interval between batch writes.

- `text_embeddings` (`object`) (optional)

  Configuration for creating text embeddings for certain tables using a Vertex AI model.

  A remote model must be attached to the dataset before using this feature, and its name supplied in the `remote_model_name` field.
  
  - `remote_model_name` (`string`) (required)

    The name of the remote model to use for text embeddings.

  - `tables` (`array`) (required)

    The tables to create text embeddings for. Each one must have its own configuration.

    - `source_table_name` (`string`) (required)
      
      The name of the source table to create text embeddings for.

    - `target_table_name` (`string`) (required)
      
      The name of the target table in which to store the embeddings.

    - `embed_columns` (`array`) (required)

      The columns to use as content for the embeddings generation function. They will be concatenated in the order they are provided.

    - `metadata_columns` (`array`) (optional)
      
      Which columns to copy as-is from the source table to the target table. `_cq_id` is always included.

  - `text_splitter` (`object`) (optional)

    The text splitter configuration to use for text embeddings. Currently only `recursive_text` is supported.

    - `recursive_text` (`object`) (required)

      - `chunk_size` (`integer`) (required)

        The size of the chunks in characters (not tokens). Defaults to `1000`.

      - `chunk_overlap` (`integer`) (required)
      
        The overlap between chunks in characters (not tokens). Defaults to `100`.

## Underlying library

We use the official [cloud.google.com/go/bigquery](https://pkg.go.dev/cloud.google.com/go/bigquery) package for database connection.
