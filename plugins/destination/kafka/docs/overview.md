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

- `client_id` (`string`) (optional) (default: `cq-destination-kafka`)

  Client ID to be set for Kafka API calls.

- `verbose` (`boolean`) (optional) (default: `false`)

  If `true`, the plugin will log all underlying Kafka client messages to the log.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Number of records to write before starting a new object.

- `topic_details` ([topic_details](#topic_details)) (optional)

  Optional parameters to set topic details.


### format_spec

- `delimiter` (`string`) (optional) (default: `,`)

  Character that will be used as want to use as the delimiter if the format type is `csv`.

- `skip_header` (`boolean`) (optional) (default: `false`)

  Specifies if the first line of a file should be the headers (when format is `csv`).


### topic_details

- `num_partitions` (`integer`) (optional) (default: `1`)

  Number of partitions for the newly created topic.

- `replication_factor` (`integer`) (optional) (default: `1`)

  Replication factor for the topic.

