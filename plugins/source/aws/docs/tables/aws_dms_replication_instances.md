# Table: aws_dms_replication_instances

https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html

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
|replication_instance|JSON|
|tags|JSON|