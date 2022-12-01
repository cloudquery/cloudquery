# Table: azure_web_publishing_profiles

https://pkg.go.dev/github.com/cloudquery/cloudquery/plugins/source/azure/client/services#PublishingProfile

The primary key for this table is **_cq_id**.

## Relations
This table depends on [azure_web_apps](azure_web_apps.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|web_app_id|String|
|publish_url|String|
|user_name|String|
|user_pwd|String|