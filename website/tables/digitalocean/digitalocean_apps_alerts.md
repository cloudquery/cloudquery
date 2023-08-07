# Table: digitalocean_apps_alerts

This table shows data for DigitalOcean Apps Alerts.

https://docs.digitalocean.com/reference/api/api-reference/#operation/apps_list_alerts

The primary key for this table is **id**.

## Relations

This table depends on [digitalocean_apps](digitalocean_apps).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`uuid`|
|component_name|`utf8`|
|spec|`json`|
|emails|`list<item: utf8, nullable>`|
|slack_webhooks|`json`|
|phase|`utf8`|
|progress|`json`|