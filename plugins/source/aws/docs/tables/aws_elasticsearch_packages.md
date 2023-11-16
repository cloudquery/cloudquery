# Table: aws_elasticsearch_packages

This table shows data for Elasticsearch Packages.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_PackageDetails.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|available_package_version|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|error_details|`json`|
|last_updated_at|`timestamp[us, tz=UTC]`|
|package_description|`utf8`|
|package_id|`utf8`|
|package_name|`utf8`|
|package_status|`utf8`|
|package_type|`utf8`|