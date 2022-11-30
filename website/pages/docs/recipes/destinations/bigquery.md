# BigQuery Destination Plugin Recipes

Full spec options for the BigQuery destination plugin are available [here](/docs/plugins/destinations/bigquery/overview#bigquery-spec).

## Basic

```yaml copy
kind: destination
spec:
  name: bigquery
  path: cloudquery/bigquery
  version: "VERSION_DESTINATION_BIGQUERY"
  write_mode: "overwrite-delete-stale"
  spec:
    project_id: ${PROJECT_ID}
    dataset_id: ${DATASET_ID}
```
