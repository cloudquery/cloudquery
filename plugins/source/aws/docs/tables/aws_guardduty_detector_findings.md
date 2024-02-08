# Table: aws_guardduty_detector_findings

This table shows data for Amazon GuardDuty Detector Findings.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Finding.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **detector_arn**, **arn**).
## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|detector_arn|`utf8`|
|account_id|`utf8`|
|arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|id|`utf8`|
|region|`utf8`|
|resource|`json`|
|schema_version|`utf8`|
|severity|`float64`|
|type|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|confidence|`float64`|
|description|`utf8`|
|partition|`utf8`|
|service|`json`|
|title|`utf8`|