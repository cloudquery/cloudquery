
# Table: azure_ad_service_principal_password_credentials
PasswordCredential active Directory Password Credential information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_principal_cq_id|uuid|Unique ID of azure_ad_service_principals table (FK)|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|start_date_time|timestamp without time zone||
|end_date_time|timestamp without time zone||
|key_id|text|Key ID|
|key_value|text|Key value|
|custom_key_identifier|bytea|Custom Key Identifier|
