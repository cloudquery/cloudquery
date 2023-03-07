# Table: aws_elasticsearch_packages

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_PackageDetails.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|id (PK)|String|
|available_package_version|String|
|created_at|Timestamp|
|error_details|JSON|
|last_updated_at|Timestamp|
|package_description|String|
|package_id|String|
|package_name|String|
|package_status|String|
|package_type|String|