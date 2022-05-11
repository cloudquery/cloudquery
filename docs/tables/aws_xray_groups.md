
# Table: aws_xray_groups
Details for a group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|A list of Tags that specify information about the group.|
|filter_expression|text|The filter expression defining the parameters to include traces.|
|arn|text|The ARN of the group generated based on the GroupName.|
|group_name|text|The unique case-sensitive name of the group.|
|insights_enabled|boolean|Set the InsightsEnabled value to true to enable insights or false to disable insights.|
|notifications_enabled|boolean|Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.|
