
# Table: azure_storage_blob_services
Azure storage blob service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of azure_storage_accounts table (FK)|
|default_service_version|text|DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified.|
|delete_retention_policy_enabled|boolean|Indicates whether DeleteRetentionPolicy is enabled|
|delete_retention_policy_days|integer|Indicates the number of days that the deleted item should be retained.|
|is_versioning_enabled|boolean|Versioning is enabled if set to true|
|automatic_snapshot_policy_enabled|boolean|Deprecated in favor of isVersioningEnabled property|
|change_feed_enabled|boolean|Indicates whether change feed event logging is enabled for the Blob service|
|change_feed_retention_in_days|integer|Indicates the duration of changeFeed retention in days.|
|restore_policy_enabled|boolean|Blob restore is enabled if set to true|
|restore_policy_days|integer|how long this blob can be restored It should be great than zero and less than DeleteRetentionPolicydays|
|restore_policy_last_enabled_time|timestamp without time zone||
|restore_policy_min_restore_time|timestamp without time zone||
|container_delete_retention_policy_enabled|boolean|Indicates whether DeleteRetentionPolicy is enabled|
|container_delete_retention_policy_days|integer|Indicates the number of days that the deleted item should be retained.|
|last_access_time_tracking_policy_enable|boolean|When set to true last access time based tracking is enabled|
|last_access_time_tracking_policy_name|text|Name of the policy.|
|last_access_time_tracking_policy_tracking_granularity_in_days|integer|The field specifies blob object tracking granularity in days.|
|last_access_time_tracking_policy_blob_type|text[]|An array of predefined supported blob types.|
|sku_name|text|Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'|
|sku_tier|text|Possible values include: 'Standard', 'Premium'|
|id|text|Fully qualified resource ID.|
|name|text|The name of the resource|
|type|text|The type of the resource.|
