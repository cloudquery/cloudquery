# Table: aws_cloudhsmv2_clusters

https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|backup_policy|String|
|backup_retention_policy|JSON|
|certificates|JSON|
|cluster_id|String|
|create_timestamp|Timestamp|
|hsm_type|String|
|hsms|JSON|
|pre_co_password|String|
|security_group|String|
|source_backup_id|String|
|state|String|
|state_message|String|
|subnet_mapping|JSON|
|tag_list|JSON|
|vpc_id|String|