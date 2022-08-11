
# Table: aws_iot_policies
The output from the GetPolicy operation.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|Tags of the resource|
|creation_date|timestamp without time zone|The date the policy was created.|
|default_version_id|text|The default policy version ID.|
|generation_id|text|The generation ID of the policy.|
|last_modified_date|timestamp without time zone|The date the policy was last modified.|
|arn|text|The policy ARN.|
|document|text|The JSON document that describes the policy.|
|name|text|The policy name.|
