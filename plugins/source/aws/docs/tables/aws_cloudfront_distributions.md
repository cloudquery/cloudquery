# Table: aws_cloudfront_distributions

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|tags|JSON|
|arn (PK)|String|
|distribution_config|JSON|
|domain_name|String|
|id|String|
|in_progress_invalidation_batches|Int|
|last_modified_time|Timestamp|
|status|String|
|active_trusted_key_groups|JSON|
|active_trusted_signers|JSON|
|alias_icp_recordals|JSON|