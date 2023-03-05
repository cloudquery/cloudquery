# Table: aws_elasticsearch_versions

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListVersions.html

The composite primary key for this table is (**account_id**, **region**, **version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|version (PK)|String|
|instance_types|JSON|