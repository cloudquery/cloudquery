
# Table: aws_iot_thing_types
The definition of the thing type, including thing type name and description.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|Tags of the resource|
|arn|text|The thing type ARN.|
|creation_date|timestamp without time zone|The date and time when the thing type was created.|
|deprecated|boolean|Whether the thing type is deprecated|
|deprecation_date|timestamp without time zone|The date and time when the thing type was deprecated.|
|name|text|The name of the thing type.|
|searchable_attributes|text[]|A list of searchable thing attribute names.|
|description|text|The description of the thing type.|
