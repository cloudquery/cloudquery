
# Table: digitalocean_accounts
Account represents a DigitalOcean Account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|droplet_limit|bigint|The total number of Droplets current user or team may have active at one time.|
|floating_ip_limit|bigint|The total number of Floating IPs the current user or team may have.|
|volume_limit|bigint|The total number of volumes the current user or team may have.|
|email|text|The email address used by the current user to register for DigitalOcean.|
|uuid|text|The unique universal identifier for the current user.|
|email_verified|boolean|If true, the user has verified their account via email. False otherwise.|
|status|text|This value is one of "active", "warning" or "locked".|
|status_message|text|A human-readable message giving more details about the status of the account.|
