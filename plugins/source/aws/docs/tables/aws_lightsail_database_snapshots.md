# Table: aws_lightsail_database_snapshots



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|created_at|Timestamp|
|engine|String|
|engine_version|String|
|from_relational_database_arn|String|
|from_relational_database_blueprint_id|String|
|from_relational_database_bundle_id|String|
|from_relational_database_name|String|
|location|JSON|
|name|String|
|resource_type|String|
|size_in_gb|Int|
|state|String|
|support_code|String|