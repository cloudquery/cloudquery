```yaml copy
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_ec2_instances"]
  destinations: ["DESTINATION_NAME"]
  spec:
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
    # scheduler: dfs # options are: dfs, round-robin or shuffle
```
