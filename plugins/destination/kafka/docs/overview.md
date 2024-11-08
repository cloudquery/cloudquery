---
name: Kafka
title: Kafka Destination Plugin
description: CloudQuery Kafka destination plugin documentation
---
# Kafka Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to Kafka in various formats such as CSV, JSON. Each table will be pushed to a separate topic.

## Example

:configuration

## Plugin Spec

This is the (nested) plugin spec

- `brokers` (`[]string`) (required)

  List of brokers to connect to.

  Example broker address:

  - `"localhost:9092"` default URL for a local Kafka broker

- `format` (`string`) (required)

  Format of the output file. Supported values are `csv`, `json` and `parquet`.

- `format_spec` ([format_spec](#format_spec)) (optional)

  Optional parameters to change the format of the file.

- `compression` (`string`) (optional) (default: empty)

  Compression algorithm to use. Supported values are empty or `gzip`. Not supported for `parquet` format.

- `sasl_username` (`string`) (optional) (default: empty)

  If connecting via SASL/PLAIN, the username to use.

- `sasl_password` (`string`) (optional) (default: empty)

  If connecting via SASL/PLAIN, the password to use.

- `verbose` (`boolean`) (optional) (default: `false`)

  If `true`, the plugin will log all underlying Kafka client messages to the log.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Number of records to write before starting a new object.

- `topic_details` ([topic_details](#topic_details)) (optional)

  Optional parameters to set topic details.


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

### topic_details

- `num_partitions` (`integer`) (optional) (default: `1`)

  Number of partitions for the newly created topic.

- `replication_factor` (`integer`) (optional) (default: `1`)

  Replication factor for the topic.

