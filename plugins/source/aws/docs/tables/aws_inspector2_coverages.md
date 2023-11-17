# Table: aws_inspector2_coverages

This table shows data for Inspector2 Coverages.

https://docs.aws.amazon.com/inspector/v2/APIReference/API_ListCoverage.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|account_id|`utf8`|
|resource_id|`utf8`|
|resource_type|`utf8`|
|scan_type|`utf8`|
|last_scanned_at|`timestamp[us, tz=UTC]`|
|resource_metadata|`json`|
|scan_status|`json`|