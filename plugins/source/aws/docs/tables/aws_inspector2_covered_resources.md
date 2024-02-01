# Table: aws_inspector2_covered_resources

This table shows data for Inspector2 Covered Resources.

https://docs.aws.amazon.com/inspector/v2/APIReference/API_CoveredResource.html

The `request_account_id` and `request_region` columns are added to show from where the request was made.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **account_id**, **resource_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|account_id|`utf8`|
|resource_id|`utf8`|
|resource_type|`utf8`|
|scan_type|`utf8`|
|last_scanned_at|`timestamp[us, tz=UTC]`|
|resource_metadata|`json`|
|scan_status|`json`|