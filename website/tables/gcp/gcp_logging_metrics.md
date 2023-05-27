# Table: gcp_logging_metrics

This table shows data for GCP Logging Metrics.

https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|filter|`utf8`|
|disabled|`bool`|
|metric_descriptor|`json`|
|value_extractor|`utf8`|
|label_extractors|`json`|
|bucket_options|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|version|`utf8`|