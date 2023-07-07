# Table: digitalocean_accounts

This table shows data for DigitalOcean Accounts.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Account

The primary key for this table is **uuid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|uuid (PK)|`utf8`|
|droplet_limit|`int64`|
|floating_ip_limit|`int64`|
|reserved_ip_limit|`int64`|
|volume_limit|`int64`|
|email|`utf8`|
|email_verified|`bool`|
|status|`utf8`|
|status_message|`utf8`|
|team|`json`|