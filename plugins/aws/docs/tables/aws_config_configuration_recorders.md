
# Table: aws_config_configuration_recorders
An object that represents the recording of configuration changes of an AWS resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|name|text|The name of the recorder.|
|recording_group_all_supported|boolean|Specifies whether AWS Config records configuration changes for every supported type of regional resource.|
|recording_group_include_global_resource_types|boolean|Specifies whether AWS Config includes all supported types of global resources (for example, IAM resources) with the resources that it records.|
|recording_group_resource_types|text[]|A comma-separated list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, AWS::EC2::Instance or AWS::CloudTrail::Trail).|
|role_arn|text|Amazon Resource Name (ARN) of the IAM role used to describe the AWS resources associated with the account.|
