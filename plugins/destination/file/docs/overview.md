---
name: File
stage: GA
title: File Destination Plugin
description: CloudQuery File destination plugin for exporting to CSV, JSON and Parquet
---
# File Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to local files in various formats. It currently supports CSV, line-delimited JSON and Parquet.

This plugin is useful in local environments, but also in production environments where scalability, performance and cost are requirements. For example, this plugin can be used as part of a system that syncs sources across multiple virtual machines, uploads Parquet files to a remote storage (such as S3 or GCS), and finally loads them to data lakes such as BigQuery or Athena in batch mode. If this is your end goal, you may also want to look at more specific destination cloud storage destination plugins such as [S3](/docs/plugins/destinations/s3/overview), [GCS](/docs/plugins/destinations/gcs/overview) or [Azure Blob Storage](/docs/plugins/destinations/azblob/overview).

## Example

:configuration

## File Spec

This is the (nested) spec used by the file destination Plugin.

- `path` (`string`) (**required**)

  Path template string that determines where files will be written, for example `path/to/files/{{TABLE}}/{{UUID}}.parquet`.

  The path supports the following placeholder variables:

  - `{{TABLE}}` will be replaced with the table name
  - `{{FORMAT}}` will be replaced with the file format, such as `csv`, `json` or `parquet`. If compression is enabled, the format will be `csv.gz`, `json.gz` etc.
  - `{{UUID}}` will be replaced with a random UUID to uniquely identify each file
  - `{{YEAR}}` will be replaced with the current year in `YYYY` format
  - `{{MONTH}}` will be replaced with the current month in `MM` format
  - `{{DAY}}` will be replaced with the current day in `DD` format
  - `{{HOUR}}` will be replaced with the current hour in `HH` format
  - `{{MINUTE}}` will be replaced with the current minute in `mm` format

  **Note** that timestamps are in `UTC` and will be the current time at the time the file is written, not when the sync started.

- `format` (`string`) (**required**)

  Format of the output file.  Supported values are `csv`, `json` and `parquet`.

- `format_spec` ([format_spec](#format_spec)) (optional)

  Optional parameters to change the format of the file.

- `no_rotate` (`boolean`) (optional) (default: `false`)

  If set to `true`, the plugin will write to one file per table.
  Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.

- `compression` (`string`) (optional) (default: `""`)

  Compression algorithm to use. Supported values are `""` and `gzip`. Not supported for `parquet` format.

- `batch_size` (`integer`) (optional) (default: `10000`)

  Number of records to write before starting a new file.

- `batch_size_bytes` (`integer`) (optional) (default: `52428800` (50 MiB))

  Number of bytes (as Arrow buffer size) to write before starting a new file.

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