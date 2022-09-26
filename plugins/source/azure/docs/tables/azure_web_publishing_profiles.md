# Table: azure_web_publishing_profiles


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`azure_web_apps`](azure_web_apps.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|web_app_id|UUID|
|publish_url|String|
|user_name|String|
|user_pwd|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|