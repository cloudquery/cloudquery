
# Table: gcp_storage_metrics
storage metrics collecting by cloud monitoring service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|bucket_name|text|Name of the bucket metric is associated with|
|acl_operation_count|bigint|Usage of ACL operations count in 24 hour period|
|object_count|bigint|Total number of objects per bucket, grouped by storage class. This value is measured once per day, and there might be a delay after measuring before the value becomes available in Cloud Monitoring.|
|total_size|bigint|Total size of all objects in the bucket (in bytes), grouped by storage class. This value is measured once per day, and there might be a delay after measuring before the value becomes available in Cloud Monitoring.|
