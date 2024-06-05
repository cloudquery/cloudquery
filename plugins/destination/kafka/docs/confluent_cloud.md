## Connecting to Confluent Cloud

![Confluent Cloud Logo](../../docs/assets/confluent-logotype-white.png)

Confluent Cloud is a fully managed data streaming platform that you can use as a destination with this plugin. You can get started with a time-limited trial at [confluent.io](https://www.confluent.io/confluent-cloud/tryfree/?utm_campaign=tm.pmm_cd.cwc_partner_cloudquery_tryfree&utm_source=cloudquery&utm_medium=partnerref).

To configure CloudQuery Kafka plugin, you need to create an API key in the [Confluent Cloud Console](https://docs.confluent.io/cloud/current/access-management/authenticate/api-keys/api-keys.html#create-a-resource-api-key).

Download the file with the API key and secret and use them for the `sasl_username` and `sasl_password` properties in the Kafka plugin documentation. The file will also contain the URL for the bootstrap server. Use that in the `brokers` property in the configuration:

```yaml
kind: destination
spec:
  name: "kafka"
  path: "cloudquery/kafka"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_KAFKA"
  write_mode: "append"
  spec:
    # required - list of brokers to connect to
    brokers: ["${CONFLUENT_BOOTSTRAP_SERVER}"]
    sasl_username: "${CONFLUENT_KEY}"
    sasl_password: "${CONFLUENT_SECRET}"
    format: "json" # options: parquet, json, csv
    format_spec:
      # CSV-specific parameters:
      # delimiter: ","
      # skip_header: false

    # Optional parameters
    # compression: "" # options: gzip
    # verbose: false
    # batch_size: 1000
    topic_details:
      num_partitions: 1
      replication_factor: 1
```

### Creating a scoped key with granular access at Confluent Cloud

If you need to limit the access of CloudQuery Kafka plugin, you can create an API key with granular access. To do this, you need to set the following ACLs on the API key:

#### Cluster ACLs

| Operation | Permission  |
|------------|------------|
| `CREATE`   | `ALLOW`    |
| `DESCRIBE` | `ALLOW`    |

#### Topic ACLs

For topic name, use the prefixes of the tables from the selected source. The table below specifies permissions for topics created by the AWS source plugin:

| Topic name | Pattern type | Operation  | Permission |
|------------|--------------|------------|------------|
| `aws_`     | `PREFIXED`   | `WRITE`    | `ALLOW`    |
| `aws_`     | `PREFIXED`   | `CREATE`   | `ALLOW`    |
| `aws_`     | `PREFIXED`   | `DESCRIBE` | `ALLOW`    |

