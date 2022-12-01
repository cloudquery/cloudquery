# Table: gcp_logging_metrics



The primary key for this table is **name**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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
|version|String|