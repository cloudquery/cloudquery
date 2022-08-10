
# Table: cloudflare_account_member_roles
AccountRole defines the roles that a member can have attached.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_member_cq_id|uuid|Unique CloudQuery ID of cloudflare_account_members table (FK)|
|account_id|text|The Account ID of the resource.|
|id|text|The unique universal identifier for a Cloudflare account member role.|
|name|text|Cloudflare account member role name.|
|description|text|Cloudflare account member role description.|
|permissions|jsonb|Cloudflare account member role permissions.|
