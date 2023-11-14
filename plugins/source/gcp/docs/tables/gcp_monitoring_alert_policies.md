# Table: gcp_monitoring_alert_policies

This table shows data for GCP Monitoring Alert Policies.

https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies#AlertPolicy

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|documentation|`json`|
|user_labels|`json`|
|conditions|`json`|
|combiner|`utf8`|
|enabled|`json`|
|validity|`json`|
|notification_channels|`list<item: utf8, nullable>`|
|creation_record|`json`|
|mutation_record|`json`|
|alert_strategy|`json`|