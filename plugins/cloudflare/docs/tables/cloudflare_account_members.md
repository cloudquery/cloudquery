
# Table: cloudflare_account_members
AccountMember is the definition of a member of an account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of cloudflare_accounts table (FK)|
|account_id|text|The Account ID of the resource.|
|id|text|The unique universal identifier for a Cloudflare account member.|
|code|text||
|user_id|text|Cloudflare user id.|
|user_first_name|text|Cloudflare user first name.|
|user_last_name|text|Cloudflare user last name.|
|user_email|text|Cloudflare user email.|
|user_two_factor_authentication_enabled|boolean|True if user has enabled 2fa authentication.|
|status|text|Cloudflare account member status.|
