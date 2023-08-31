# Table: aws_elasticsearch_versions

This table shows data for Elasticsearch Versions.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListVersions.html

The composite primary key for this table is (**account_id**, **region**, **version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|version (PK)|`utf8`|
|instance_types|`json`|