
# Table: aws_iot_thing_groups
Groups allow you to manage several things at once by categorizing them into groups
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|things_in_group|text[]|Lists the things in the specified group|
|policies|text[]|Policies associated with the thing group|
|tags|jsonb|Tags of the resource|
|index_name|text|The dynamic thing group index name.|
|query_string|text|The dynamic thing group search query string.|
|query_version|text|The dynamic thing group query version.|
|status|text|The dynamic thing group status.|
|arn|text|The thing group ARN.|
|id|text|The thing group ID.|
|creation_date|timestamp without time zone|The UNIX timestamp of when the thing group was created.|
|parent_group_name|text|The parent thing group name.|
|root_to_parent_thing_groups|jsonb|The root parent thing group.|
|name|text|The name of the thing group.|
|attribute_payload_attributes|jsonb|A JSON string containing up to three key-value pair in JSON format|
|attribute_payload_merge|boolean|Specifies whether the list of attributes provided in the AttributePayload is merged with the attributes stored in the registry, instead of overwriting them. To remove an attribute, call UpdateThing with an empty attribute value|
|thing_group_description|text|The thing group description.|
|version|bigint|The version of the thing group.|
