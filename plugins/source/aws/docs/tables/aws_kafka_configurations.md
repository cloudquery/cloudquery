# Table: aws_kafka_configurations

https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-configuration.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|creation_time|Timestamp|
|description|String|
|kafka_versions|StringArray|
|latest_revision|JSON|
|name|String|
|state|String|