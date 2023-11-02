# Table: aws_codeartifact_domains

This table shows data for AWS CodeArtifact Domains.

https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|asset_size_bytes|`int64`|
|created_time|`timestamp[us, tz=UTC]`|
|encryption_key|`utf8`|
|name|`utf8`|
|owner|`utf8`|
|repository_count|`int64`|
|s3_bucket_arn|`utf8`|
|status|`utf8`|