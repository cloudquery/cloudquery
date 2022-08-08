
# Table: aws_guardduty_detector_members
Contains information about the member account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|detector_cq_id|uuid|Unique CloudQuery ID of aws_guardduty_detectors table (FK)|
|account_id|text|The ID of the member account.|
|email|text|The email address of the member account.|
|master_id|text|The administrator account ID.|
|relationship_status|text|The status of the relationship between the member and the administrator.|
|updated_at|timestamp without time zone|The last-updated timestamp of the member.|
|detector_id|text|The detector ID of the member account.|
|invited_at|timestamp without time zone|The timestamp when the invitation was sent.|
