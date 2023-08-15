# Table: aws_guardduty_detector_intel_sets

This table shows data for Amazon GuardDuty Detector Intel Sets.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetThreatIntelSet.html

The composite primary key for this table is (**detector_arn**, **name**).

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|detector_arn (PK)|`utf8`|
|format|`utf8`|
|location|`utf8`|
|name (PK)|`utf8`|
|status|`utf8`|
|tags|`json`|