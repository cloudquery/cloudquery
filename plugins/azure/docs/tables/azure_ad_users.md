
# Table: azure_ad_users
User active Directory user information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|immutable_id|text|This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account It is used to associate an on-premises Active Directory user account with their Azure AD user object|
|usage_location|text|A two letter country code (ISO standard 3166) Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries Examples include: "US", "JP", and "GB"|
|given_name|text|The given name for the user|
|surname|text|The user's surname (family name or last name)|
|user_type|text|A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest' Possible values include: 'Member', 'Guest'|
|account_enabled|boolean|Whether the account is enabled|
|display_name|text|The display name of the user|
|user_principal_name|text|The principal name of the user|
|mail_nickname|text|The mail alias for the user|
|mail|text|The primary email address of the user|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|object_id|text|The object ID|
|deletion_timestamp_time|timestamp without time zone||
|object_type|text|Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'|
