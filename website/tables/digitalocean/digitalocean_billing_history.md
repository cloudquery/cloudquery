# Table: digitalocean_billing_history

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|description|String|
|amount|String|
|invoice_id|String|
|invoice_uuid|String|
|date|Timestamp|
|type|String|