# Table: gcp_monitoring_alert_policies


The primary key for this table is **name**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name (PK)|String|
|display_name|String|
|documentation|JSON|
|user_labels|JSON|
|conditions|JSON|
|combiner|Int|
|enabled|JSON|
|validity|JSON|
|notification_channels|StringArray|
|creation_record|JSON|
|mutation_record|JSON|
|alert_strategy|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|