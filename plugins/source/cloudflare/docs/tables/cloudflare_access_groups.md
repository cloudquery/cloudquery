
# Table: cloudflare_access_groups
AccessGroup defines a group for allowing or disallowing access to one or more Access applications.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|zone_id|text|Zone identifier tag.|
|id|text|The unique identifier for the Access group.|
|created_at|timestamp without time zone|Hashed script content, can be used in a If-None-Match header when updating.|
|updated_at|timestamp without time zone|Size of the script, in bytes.|
|name|text|The name of the Access group.|
|include|jsonb|Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.|
|exclude|jsonb|Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.|
|require|jsonb|Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.|
