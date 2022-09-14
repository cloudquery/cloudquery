
# Table: digitalocean_space_acls
 list of elements describing allowed methods for a specific origin.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|space_cq_id|uuid|Unique CloudQuery ID of digitalocean_spaces table (FK)|
|space_name|text|name of the space.|
|permission|text|The level of access granted. At this time, the only supported values are FULL_CONTROL and READ.|
|type|text|Type of grantee|
|display_name|text|Screen name of the grantee.|
|email_address|text|Email address of the grantee|
|grantee_id|text|The canonical user ID of the grantee.|
|uri|text|A URI specifying a group of users. At this time, only http://acs.amazonaws.com/groups/global/AllUsers is supported.|
