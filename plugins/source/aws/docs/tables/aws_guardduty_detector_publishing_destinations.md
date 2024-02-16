# Table: aws_guardduty_detector_publishing_destinations

This table shows data for Amazon GuardDuty Detector Publishing Destinations.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DescribePublishingDestination.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **detector_arn**, **arn**, **destination_id**).
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
|destination_id|`utf8`|
|destination_type|`utf8`|
|status|`utf8`|