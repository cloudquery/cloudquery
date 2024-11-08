---
name: Azure Blob Storage
stage: GA
title: Azure Blob Destination Plugin
description: CloudQuery Azure Blob destination plugin documentation
---
# Azure Blob Storage Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to remote Azure Blob Storage storage in various formats such as CSV, JSON and Parquet.

## Authentication

:authentication

## Example

This example configures an Azure blob storage destination, to create CSV files in `https://cqdestinationazblob.blob.core.windows.net/test/path/to/files`.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

:configuration

The Azure Blob destination utilizes batching, and supports `batch_size`, `batch_size_bytes` and `batch_timeout` options (see below).

## Azure Blob Spec

This is the (nested) spec used by the Azure blob destination Plugin.

- `storage_account` (`string`) (required)

  Storage account where to sync the files.

- `container` (`string`) (required)

  Storage container inside the storage account where to sync the files.

- `path` (`string`) (required)

  Path to where the files will be uploaded in the above bucket.

- `no_rotate` (`boolean`) (optional) (default: `false`)

  If set to `true`, the plugin will write to one file per table.
  Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.

- `format` (`string`) (required)

  Format of the output file. Supported values are `csv`, `json` and `parquet`.

- `format_spec` ([format_spec](#format_spec)) (optional)

  Optional parameters to change the format of the file.

- `compression` (`string`) (optional) (default: empty)

  Compression algorithm to use. Supported values are empty or `gzip`. Not supported for `parquet` format.

- `batch_size` (`integer`) (optional) (default: `10000`)

  Number of records to write before starting a new object.

- `batch_size_bytes` (`integer`) (optional) (default: `52428800` (50 MiB))

  Number of bytes (as Arrow buffer size) to write before starting a new object.

- `batch_timeout` (`duration`) (optional) (default: `30s` (30 seconds))

  Maximum interval between batch writes.

### format_spec

#### CSV

- `delimiter` (`string`) (optional) (default: `,`)

  Delimiter to use in the CSV file.

- `skip_header` (`boolean`) (optional) (default: `false`)

  If set to `true`, the CSV file will not contain a header row as the first row.

#### JSON

Reserved for future use.

#### Parquet

- `version` (`string`) (optional) (default: `v2Latest`)

  Parquet format version to use. Supported values are `v1.0`, `v2.4`, `v2.6` and `v2Latest`.
  `v2Latest` is an alias for the latest version available in the Parquet library which is currently `v2.6`.

  Useful when the reader consuming the Parquet files does not support the latest version.

- `root_repetition` (`string`) (optional) (default: `repeated`)

  [Repetition option to use for the root node](https://github.com/apache/arrow/issues/20243). Supported values are `undefined`, `required`, `optional` and `repeated`.

  Some Parquet readers require a specific root repetition option to be able to read the file. For example, importing Parquet files into [Snowflake](https://www.snowflake.com/en/) requires the root repetition to be `undefined`.

- `max_row_group_length` (`integer`) (optional) (default: `134217728` (= 128 * 1024 * 1024))

  The maximum number of rows in a single row group. Use a lower number to reduce memory usage when reading the Parquet files, and a higher number to increase the efficiency of reading the Parquet files.