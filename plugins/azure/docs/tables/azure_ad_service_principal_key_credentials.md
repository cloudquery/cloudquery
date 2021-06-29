
# Table: azure_ad_service_principal_key_credentials
KeyCredential active Directory Key Credential information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_principal_id|uuid|Unique ID of azure_ad_service_principals table (FK)|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|start_date|timestamp without time zone|Start date.|
|end_date|timestamp without time zone|End date.|
|key_value|text|Key value|
|key_id|text|Key ID|
|usage|text|Usage Acceptable values are 'Verify' and 'Sign'|
|key_type|text|Type Acceptable values are 'AsymmetricX509Cert' and 'Symmetric'|
|custom_key_identifier|text|Custom Key Identifier|
