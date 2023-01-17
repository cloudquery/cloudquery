# Table: azure_subscription_subscription_locations

The primary key for this table is **id**.

## Relations

This table depends on [azure_subscription_subscriptions](azure_subscription_subscriptions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|display_name|String|
|latitude|String|
|longitude|String|
|name|String|
|subscription_id|String|