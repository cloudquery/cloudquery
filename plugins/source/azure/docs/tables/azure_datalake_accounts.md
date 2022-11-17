# Table: azure_datalake_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics#Account

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|name|String|
|firewall_allow_azure_ips|String|
|firewall_state|String|
|max_degree_of_parallelism|Int|
|max_degree_of_parallelism_per_job|Int|
|max_job_count|Int|
|new_tier|String|
|public_data_lake_store_accounts|JSON|
|query_store_retention|Int|
|account_id|String|
|compute_policies|JSON|
|creation_time|Timestamp|
|current_tier|String|
|data_lake_store_accounts|JSON|
|debug_data_access_level|String|
|default_data_lake_store_account|String|
|default_data_lake_store_account_type|String|
|endpoint|String|
|firewall_rules|JSON|
|hive_metastores|JSON|
|last_modified_time|Timestamp|
|max_active_job_count_per_user|Int|
|max_job_running_time_in_min|Int|
|max_queued_job_count_per_user|Int|
|min_priority_per_job|Int|
|provisioning_state|String|
|state|String|
|storage_accounts|JSON|
|system_max_degree_of_parallelism|Int|
|system_max_job_count|Int|
|virtual_network_rules|JSON|
|tags|JSON|
|type|String|