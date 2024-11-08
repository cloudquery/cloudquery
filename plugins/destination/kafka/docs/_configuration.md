This example configures connects to a Kafka destination using SASL plain authentication and pushes messages in JSON format.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml copy
kind: destination
spec:
  name: "kafka"
  path: "cloudquery/kafka"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_KAFKA"
  write_mode: "append"
  spec:
    # required - list of brokers to connect to
    brokers: ["<broker-host>:<broker-port>"]
    # optional - if connecting via SASL/PLAIN, the username and password to use. If not set, no authentication will be used.
    sasl_username: "${KAFKA_SASL_USERNAME}"
    sasl_password: "${KAFKA_SASL_PASSWORD}"
    format: "json" # options: parquet, json, csv
    format_spec:
      # CSV specific parameters:
      # delimiter: ","
      # skip_header: false
      # Parquet specific parameters:
      # version: "v2Latest"
      # root_repetition: "repeated"
      # max_row_group_length: 134217728 # 128 * 1024 * 1024

    # Optional parameters
    # compression: "" # options: gzip
    # verbose: false
    # batch_size: 1000
    # topic_details:
      # num_partitions: 1
      # replication_factor: 1
```

Note that the Kafka plugin only supports `append` `write_mode`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).
