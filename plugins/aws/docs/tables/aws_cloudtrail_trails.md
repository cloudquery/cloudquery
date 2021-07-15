
# Table: aws_cloudtrail_trails
The settings for a trail.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|cloudwatch_logs_log_group_name|text||
|is_logging|boolean| Whether the CloudTrail is currently logging AWS API calls.|
|latest_cloud_watch_logs_delivery_error|text|Displays any CloudWatch Logs error that CloudTrail encountered when attempting to deliver logs to CloudWatch Logs.|
|latest_cloud_watch_logs_delivery_time|timestamp without time zone|Displays the most recent date and time when CloudTrail delivered logs to CloudWatch Logs.|
|latest_delivery_error|text|Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver log files to the designated bucket.|
|latest_delivery_time|timestamp without time zone|Specifies the date and time that CloudTrail last delivered log files to an account's Amazon S3 bucket.|
|latest_digest_delivery_error|text|Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver a digest file to the designated bucket.|
|latest_digest_delivery_time|timestamp without time zone|Specifies the date and time that CloudTrail last delivered a digest file to an account's Amazon S3 bucket.|
|latest_notification_error|text| Displays any Amazon SNS error that CloudTrail encountered when attempting to send a notification.|
|latest_notification_time|timestamp without time zone|Specifies the date and time of the most recent Amazon SNS notification that CloudTrail has written a new log file to an account's Amazon S3 bucket.|
|start_logging_time|timestamp without time zone|Specifies the most recent date and time when CloudTrail started recording API calls for an AWS account.|
|stop_logging_time|timestamp without time zone|Specifies the most recent date and time when CloudTrail stopped recording API calls for an AWS account.|
|cloud_watch_logs_log_group_arn|text|Specifies an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered.|
|cloud_watch_logs_role_arn|text|Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.|
|has_custom_event_selectors|boolean|Specifies if the trail has custom event selectors.|
|has_insight_selectors|boolean|Specifies whether a trail has insight types specified in an InsightSelector list.|
|home_region|text|The region in which the trail was created.|
|include_global_service_events|boolean|Set to True to include AWS API calls from AWS global services such as IAM.|
|is_multi_region_trail|boolean|Specifies whether the trail exists only in one region or exists in all regions.|
|is_organization_trail|boolean|Specifies whether the trail is an organization trail.|
|kms_key_id|text|Specifies the KMS key ID that encrypts the logs delivered by CloudTrail.|
|log_file_validation_enabled|boolean|Specifies whether log file validation is enabled.|
|name|text|Name of the trail set by calling CreateTrail.|
|s3_bucket_name|text|Name of the Amazon S3 bucket into which CloudTrail delivers your trail files.|
|s3_key_prefix|text|Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.|
|sns_topic_arn|text|Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send notifications when log files are delivered.|
|sns_topic_name|text|This field is no longer in use.|
|arn|text|Specifies the ARN of the trail.|
