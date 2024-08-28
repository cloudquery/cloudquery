```yaml copy
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "v27.0.0"
  tables: ["aws_ec2_instances"]
  destinations: ["DESTINATION_NAME"]
  # Learn more about the configuration options at https://cql.ink/aws_source
  spec:
    concurrency: 100
    # Optional parameters
    # regions: []
    # accounts: []
    # org: nil
    # concurrency: 50000
    # initialization_concurrency: 4
    # aws_debug: false
    # max_retries: 10
    # max_backoff: 30
    # custom_endpoint_url: ""
    # custom_endpoint_hostname_immutable: nil # required when custom_endpoint_url is set
    # custom_endpoint_partition_id: "" # required when custom_endpoint_url is set
    # custom_endpoint_signing_region: "" # required when custom_endpoint_url is set
    # use_paid_apis: false
    # table_options: nil
    # scheduler: shuffle # options are: dfs, round-robin or shuffle
    # use_nested_table_rate_limiting: false
    # enable_api_level_tracing: false
```
