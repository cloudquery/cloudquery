
# Table: aws_iot_things
The properties of the thing, including thing name, thing type name, and a list of thing attributes.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|principals|text[]|Principals associated with the thing|
|attributes|jsonb|A list of thing attributes which are name-value pairs.|
|arn|text|The thing ARN.|
|name|text|The name of the thing.|
|type_name|text|The name of the thing type, if the thing has been associated with a type.|
|version|bigint|The version of the thing record in the registry.|
