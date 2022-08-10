
# Table: aws_shield_subscriptions
Information about the Shield Advanced subscription for an account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|protection_group_limits_max_protection_groups|integer|The maximum number of protection groups that you can have at one time|
|protection_group_limits_arbitrary_pattern_limits_max_members|integer|The maximum number of resources you can specify for a single arbitrary pattern in a protection group|
|protected_resource_type_limits|jsonb|The maximum number of resource types that you can specify in a protection|
|auto_renew|text|If ENABLED, the subscription will be automatically renewed at the end of the existing subscription period|
|end_time|timestamp without time zone|The date and time your subscription will end|
|limits|jsonb|Specifies how many protections of a given type you can create|
|proactive_engagement_status|text|If ENABLED, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support|
|start_time|timestamp without time zone|The start time of the subscription, in Unix time in seconds|
|arn|text|The ARN (Amazon Resource Name) of the subscription|
|time_commitment_in_seconds|integer|The length, in seconds, of the Shield Advanced subscription for the account|
