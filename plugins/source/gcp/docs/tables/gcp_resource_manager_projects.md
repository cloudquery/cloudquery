
# Table: gcp_resource_manager_projects
A project is a high-level Google Cloud entity It is a container for ACLs, APIs, App Engine Apps, VMs, and other Google Cloud Platform resources
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|policy|jsonb|Access control policy for a resource|
|create_time|timestamp without time zone|Creation time|
|delete_time|timestamp without time zone|The time at which this resource was requested for deletion|
|display_name|text|A user-assigned display name of the project When present it must be between 4 to 30 characters Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point|
|etag|text|A checksum computed by the server based on the current value of the Project resource This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding|
|labels|jsonb|The labels associated with this project|
|name|text|The unique resource name of the project It is an int64 generated number prefixed by "projects/"|
|parent|text|A reference to a parent Resource eg, `organizations/123` or `folders/876`|
|project_id|text|Immutable The unique, user-assigned id of the project It must be 6 to 30 lowercase ASCII letters, digits, or hyphens It must start with a letter Trailing hyphens are prohibited|
|state|text|The project lifecycle state  Possible values:   "STATE_UNSPECIFIED" - Unspecified state This is only used/useful for distinguishing unset values   "ACTIVE" - The normal and active state   "DELETE_REQUESTED" - The project has been marked for deletion by the user (by invoking DeleteProject) or by the system (Google Cloud Platform) This can generally be reversed by invoking UndeleteProject|
|update_time|timestamp without time zone|The most recent time this resource was modified|
