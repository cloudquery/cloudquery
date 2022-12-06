# Table: azure_datalake_analytics_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account#DataLakeAnalyticsAccount

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|default_data_lake_store_account|String|
|data_lake_store_accounts|JSON|
|storage_accounts|JSON|
|compute_policies|JSON|
|firewall_rules|JSON|
|firewall_state|String|
|firewall_allow_azure_ips|String|
|new_tier|String|
|current_tier|String|
|max_job_count|Int|
|system_max_job_count|Int|
|max_degree_of_parallelism|Int|
|system_max_degree_of_parallelism|Int|
|max_degree_of_parallelism_per_job|Int|
|min_priority_per_job|Int|
|query_store_retention|Int|
|account_id|UUID|
|provisioning_state|String|
|state|String|
|creation_time|Timestamp|
|last_modified_time|Timestamp|
|endpoint|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|