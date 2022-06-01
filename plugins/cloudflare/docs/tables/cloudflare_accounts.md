
# Table: cloudflare_accounts
Account represents the root object that owns resources.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|The unique universal identifier for a Cloudflare account.|
|name|text|Cloudflare account name.|
|type|text|Cloudflare account type.|
|created_on|timestamp without time zone|Creation timestamp of the account.|
|enforce_two_factor|boolean|True if the account has enforce 2fa authentication.|
