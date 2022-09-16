
# Table: aws_ssm_parameters
Metadata includes information like the ARN of the last user and the date/time the parameter was last used
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|allowed_pattern|text|A parameter name can include only the following letters and symbols a-zA-Z0-9_-|
|data_type|text|The data type of the parameter, such as text or aws:ec2:image|
|description|text|Description of the parameter actions|
|key_id|text|The ID of the query key used for this parameter|
|last_modified_date|timestamp without time zone|Date the parameter was last changed or updated|
|last_modified_user|text|Amazon Resource Name (ARN) of the Amazon Web Services user who last changed the parameter|
|name|text|The parameter name|
|tier|text|The parameter tier|
|type|text|The type of parameter|
|version|bigint|The parameter version|
