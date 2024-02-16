# Table: aws_guardduty_detector_ip_sets

This table shows data for Amazon GuardDuty Detector IP Sets.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetIPSet.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **detector_arn**, **arn**, **name**).
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
|arn|`utf8`|
|format|`utf8`|
|location|`utf8`|
|name|`utf8`|
|status|`utf8`|
|tags|`json`|
|id|`utf8`|