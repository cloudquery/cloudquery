# Table: digitalocean_monitoring_alert_policies

This table shows data for DigitalOcean Monitoring Alert Policies.

https://docs.digitalocean.com/reference/api/api-reference/#operation/monitoring_list_alertPolicy

The primary key for this table is **uuid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|uuid (PK)|`utf8`|
|type|`utf8`|
|description|`utf8`|
|compare|`utf8`|
|value|`float64`|
|window|`utf8`|
|entities|`list<item: utf8, nullable>`|
|tags|`list<item: utf8, nullable>`|
|alerts|`json`|
|enabled|`bool`|