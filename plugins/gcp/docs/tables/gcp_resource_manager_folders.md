
# Table: gcp_resource_manager_folders
A folder in an organization's resource hierarchy, used to organize that organization's resources
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|policy|jsonb|Access control policy for a resource|
|create_time|timestamp without time zone|Timestamp when the folder was created|
|delete_time|timestamp without time zone|Timestamp when the folder was requested to be deleted|
|display_name|text|The folder's display name A folder's display name must be unique amongst its siblings For example, no two folders with the same parent can share the same display name The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`|
|etag|text|A checksum computed by the server based on the current value of the folder resource This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding|
|name|text|The resource name of the folder Its format is `folders/{folder_id}`, for example: "folders/1234"|
|parent|text|The folder's parent's resource name Updates to the folder's parent must be performed using MoveFolder|
|state|text|The lifecycle state of the folder Updates to the state must be performed using DeleteFolder and UndeleteFolder  Possible values:   "STATE_UNSPECIFIED" - Unspecified state   "ACTIVE" - The normal and active state   "DELETE_REQUESTED" - The folder has been marked for deletion by the user|
|update_time|timestamp without time zone|Timestamp when the folder was last modified|
