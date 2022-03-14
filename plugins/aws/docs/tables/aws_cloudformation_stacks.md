
# Table: aws_cloudformation_stacks
The Stack data type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text||
|creation_time|timestamp without time zone|The time at which the stack was created.  This member is required.|
|stack|text|The name associated with the stack.  This member is required.|
|status|text|Current status of the stack.  This member is required.|
|capabilities|text[]|The capabilities allowed in the stack.|
|change_set_id|text|The unique ID of the change set.|
|deletion_time|timestamp without time zone|The time the stack was deleted.|
|description|text|A user-defined description associated with the stack.|
|disable_rollback|boolean|Boolean to enable or disable rollback on stack creation failures:  * true: disable rollback.  * false: enable rollback.|
|stack_drift_status|text|Status of the stack's actual configuration compared to its expected template configuration.  * DRIFTED: The stack differs from its expected template configuration|
|drift_last_check_timestamp|timestamp without time zone|Most recent time when a drift detection operation was initiated on the stack, or any of its individual resources that support drift detection.|
|enable_termination_protection|boolean|Whether termination protection is enabled for the stack|
|last_updated_time|timestamp without time zone|The time the stack was last updated|
|notification_arns|text[]|Amazon SNS topic Amazon Resource Names (ARNs) to which stack related events are published.|
|parameters|jsonb|A list of Parameter structures.|
|parent_id|text|For nested stacks--stacks created as resources for another stack--the stack ID of the direct parent of this stack|
|role_arn|text|The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that's associated with the stack|
|rollback_configuration_monitoring_time_in_minutes|integer|The amount of time, in minutes, during which CloudFormation should monitor all the rollback triggers after the stack creation or update operation deploys all necessary resources|
|rollback_configuration_rollback_triggers|jsonb|The triggers to monitor during stack creation or update actions|
|root_id|text|For nested stacks--stacks created as resources for another stack--the stack ID of the top-level stack to which the nested stack ultimately belongs|
|id|text|Unique identifier of the stack.|
|stack_status_reason|text|Success/failure message associated with the stack status.|
|tags|jsonb|A list of Tags that specify information about the stack.|
|timeout_in_minutes|integer|The amount of time within which stack creation should complete.|
