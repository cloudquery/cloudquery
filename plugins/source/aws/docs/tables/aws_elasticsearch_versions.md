# Table: aws_elasticsearch_versions

This table shows data for Elasticsearch Versions.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListVersions.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|version|`utf8`|
|instance_types|`json`|