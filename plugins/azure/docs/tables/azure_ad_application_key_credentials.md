
# Table: azure_ad_application_key_credentials
KeyCredential active Directory Key Credential information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|application_cq_id|uuid|Unique ID of azure_ad_applications table (FK)|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|start_date_time|timestamp without time zone||
|end_date_time|timestamp without time zone||
|value|text|Key value|
|key_id|text|Key ID|
|usage|text|Usage Acceptable values are 'Verify' and 'Sign'|
|type|text|Type Acceptable values are 'AsymmetricX509Cert' and 'Symmetric'|
|custom_key_identifier|text|Custom Key Identifier|
