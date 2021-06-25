
# Table: azure_ad_user_sign_in_names
SignInName contains information about a sign-in name of a local account user in an Azure Active Directory B2C tenant
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_id|uuid|Unique ID of azure_ad_users table (FK)|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|signin_type|text|A string value that can be used to classify user sign-in types in your directory, such as 'emailAddress' or 'userName'|
|signin_value|text|The sign-in used by the local account Must be unique across the company/tenant For example, 'johnc@examplecom'|
