
# Table: azure_datalake_analytics_account_compute_policies
ComputePolicy data Lake Analytics compute policy information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analytics_account_cq_id|uuid|Unique CloudQuery ID of azure_datalake_analytics_accounts table (FK)|
|object_id|uuid|The AAD object identifier for the entity to create a policy for|
|object_type|text|The type of AAD object the object identifier refers to|
|max_degree_of_parallelism_per_job|integer|The maximum degree of parallelism per job this user can use to submit jobs|
|min_priority_per_job|integer|The minimum priority per job this user can use to submit jobs|
|id|text|The resource identifier|
|name|text|The resource name|
|type|text|The resource type|
