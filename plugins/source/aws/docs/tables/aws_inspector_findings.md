
# Table: aws_inspector_findings
Contains information about an Amazon Inspector finding
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|arn|text|The ARN that specifies the finding|
|attributes|jsonb|The system-defined attributes for the finding|
|created_at|timestamp without time zone|The time when the finding was generated|
|updated_at|timestamp without time zone|The time when AddAttributesToFindings is called|
|user_attributes|jsonb|The user-defined attributes that are assigned to the finding|
|asset_attributes|jsonb|A collection of attributes of the host from which the finding is generated|
|asset_type|text|The type of the host from which the finding is generated|
|confidence|bigint|This data element is currently not used|
|description|text|The description of the finding|
|id|text|The ID of the finding|
|indicator_of_compromise|boolean|This data element is currently not used|
|numeric_severity|float|The numeric value of the finding severity|
|recommendation|text|The recommendation for the finding|
|schema_version|bigint|The schema version of this data type|
|service|text|The data element is set to "Inspector"|
|service_attributes|jsonb|This data type is used in the Finding data type|
|severity|text|The finding severity|
|title|text|The name of the finding|
