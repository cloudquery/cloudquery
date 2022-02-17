
# Table: azure_datalake_analytics_accounts
Data Lake Analytics account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|default_data_lake_store_account|text|The default Data Lake Store account associated with this account|
|firewall_state|text|The current state of the IP address firewall for this account|
|firewall_allow_azure_ips|text|The current state of allowing or disallowing IPs originating within Azure through the firewall|
|new_tier|text|The commitment tier for the next month|
|current_tier|text|The commitment tier in use for the current month|
|max_job_count|integer|The maximum supported jobs running under the account at the same time|
|system_max_job_count|integer|The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account|
|max_degree_of_parallelism|integer|The maximum supported degree of parallelism for this account|
|system_max_degree_of_parallelism|integer|The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account|
|max_degree_of_parallelism_per_job|integer|The maximum supported degree of parallelism per job for this account|
|min_priority_per_job|integer|The minimum supported priority per job for this account|
|query_store_retention|integer|The number of days that job metadata is retained|
|account_id|uuid|The unique identifier associated with this Data Lake Analytics account|
|provisioning_state|text|The provisioning status of the Data Lake Analytics account|
|state|text|The state of the Data Lake Analytics account|
|creation_time|timestamp without time zone||
|last_modified_time|timestamp without time zone||
|endpoint|text|The full CName endpoint for this account|
|id|text|The resource identifier|
|name|text|The resource name|
|type|text|The resource type|
|location|text|The resource location|
|tags|jsonb|The resource tags|
