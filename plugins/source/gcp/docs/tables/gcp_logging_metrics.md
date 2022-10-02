# Table: gcp_logging_metrics


The primary key for this table is **name**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name (PK)|String|
|description|String|
|filter|String|
|disabled|Bool|
|metric_descriptor|JSON|
|value_extractor|String|
|label_extractors|JSON|
|bucket_options|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|version|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|