
# Table: azure_ad_groups
ADGroup active Directory group information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|display_name|text|The display name of the group|
|mail_enabled|boolean|Whether the group is mail-enabled Must be false This is because only pure security groups can be created using the Graph API|
|mail_nickname|text|The mail alias for the group|
|security_enabled|boolean|Whether the group is security-enable|
|mail|text|The primary email address of the group|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|object_id|text|The object ID|
|deletion_timestamp_time|timestamp without time zone|The time at which the directory object was deleted.|
|object_type|text|Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'|
