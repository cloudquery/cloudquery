---
name: GCS
stage: GA
title: GCS Destination Plugin
description: CloudQuery GCS destination plugin documentation
---
# GCS (Google Cloud Storage) Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to remote GCS (Google Cloud Storage) storage in various formats such as CSV, JSON and Parquet.

This is useful in various use-cases, especially in data lakes where you can query the data direct from Athena or load it to various data warehouses such as BigQuery, RedShift, Snowflake and others.

## Example

:configuration

The GCS destination utilizes batching, and supports `batch_size`, `batch_size_bytes` and `batch_timeout` options (see below).

## GCS Spec

This is the (nested) spec used by the CSV destination Plugin.

- `bucket` (`string`) (required)

  Bucket where to sync the files.

- `path` (`string`) (required)

  Path to where the files will be uploaded in the above bucket, for example `path/to/files/{{TABLE}}/{{UUID}}.parquet`.

  If no path variables are present, the path will be appended with TABLE, FORMAT and Compression extension by default.

  The path supports the following placeholder variables:

  - `{{TABLE}}` will be replaced with the table name
  - `{{SYNC_ID}}` will be replaced with the unique identifier of the sync. This value is a UUID and is randomly generated for each sync.
  - `{{FORMAT}}` will be replaced with the file format, such as `csv`, `json` or `parquet`. If compression is enabled, the format will be `csv.gz`, `json.gz` etc.
  - `{{UUID}}` will be replaced with a random UUID to uniquely identify each file
  - `{{YEAR}}` will be replaced with the current year in `YYYY` format
  - `{{MONTH}}` will be replaced with the current month in `MM` format
  - `{{DAY}}` will be replaced with the current day in `DD` format
  - `{{HOUR}}` will be replaced with the current hour in `HH` format
  - `{{MINUTE}}` will be replaced with the current minute in `mm` format

- `format` (`string`) (required)

  Format of the output file. Supported values are `csv`, `json` and `parquet`.

- `format_spec` ([format_spec](#format_spec)) (optional)

  Optional parameters to change the format of the file.

- `compression` (`string`) (optional) (default: empty)

  Compression algorithm to use. Supported values are empty or `gzip`. Not supported for `parquet` format.

- `no_rotate` (`boolean`) (optional) (default: `false`)

  If set to `true`, the plugin will write to one file per table.
  Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.

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

## Authentication

:authentication