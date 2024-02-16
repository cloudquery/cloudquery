# Table: aws_cloudfront_distributions

This table shows data for Cloudfront Distributions.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|arn|`utf8`|
|distribution_config|`json`|
|domain_name|`utf8`|
|id|`utf8`|
|in_progress_invalidation_batches|`int64`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|active_trusted_key_groups|`json`|
|active_trusted_signers|`json`|
|alias_icp_recordals|`json`|