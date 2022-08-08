
# Table: aws_config_configuration_recorders
An object that represents the recording of configuration changes of an AWS resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|Amazon Resource Name (ARN) of the config recorder.|
|name|text|The name of the recorder.|
|recording_group_all_supported|boolean|Specifies whether AWS Config records configuration changes for every supported type of regional resource.|
|recording_group_include_global_resource_types|boolean|Specifies whether AWS Config includes all supported types of global resources (for example, IAM resources) with the resources that it records.|
|recording_group_resource_types|text[]|A comma-separated list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, AWS::EC2::Instance or AWS::CloudTrail::Trail).|
|role_arn|text|Amazon Resource Name (ARN) of the IAM role used to describe the AWS resources associated with the account.|
|status_last_error_code|text|The error code indicating that the recording failed.|
|status_last_error_message|text|The message indicating that the recording failed due to an error.|
|status_last_start_time|timestamp without time zone|The time the recorder was last started.|
|status_last_status|text|The last (previous) status of the recorder.|
|status_last_status_change_time|timestamp without time zone|The time when the status was last changed.|
|status_last_stop_time|timestamp without time zone|The time the recorder was last stopped.|
|status_recording|boolean|Specifies whether or not the recorder is currently recording.|
