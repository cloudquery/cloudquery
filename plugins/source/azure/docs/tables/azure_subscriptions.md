# Table: azure_subscriptions


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|authorization_source|String|
|managed_by_tenants|JSON|
|subscription_policies|JSON|
|tags|JSON|
|display_name|String|
|id (PK)|String|
|state|String|
|tenant_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|