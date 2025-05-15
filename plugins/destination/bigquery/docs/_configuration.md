```yaml copy
kind: destination
spec:
  name: bigquery
  path: cloudquery/bigquery
  registry: cloudquery
  version: "VERSION_DESTINATION_BIGQUERY"
  write_mode: "append"
  # Learn more about the configuration options at https://cql.ink/bigquery_destination
  spec:
    project_id: ${PROJECT_ID}
    dataset_id: ${DATASET_ID}
    # Optional parameters
    # dataset_location: ""
    # time_partitioning: none # options: "none", "hour", "day"
    # service_account_key_json: ""
    # endpoint: ""
    # batch_size: 10000
    # batch_size_bytes: 5242880 # 5 MiB
    # batch_timeout: 10s
    # client_project_id: "*detect-project-id*"
```

This example above expects the following environment variables to be set:

  * `PROJECT_ID` - The Google Cloud Project ID
  * `DATASET_ID` - The Google Cloud BigQuery Dataset ID

`client_project_id` variable can be used to run BigQuery queries in a project different from where the destination table is located. 
If you set client_project_id to `*detect-project-id*`, it will automatically detect the project ID from the environment variable or application default credentials.