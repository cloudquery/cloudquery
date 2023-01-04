# Table: gcp_cloudidentity_client_states

https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers.clientStates#ClientState

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|asset_tags|StringArray|
|compliance_state|String|
|create_time|String|
|custom_id|String|
|etag|String|
|health_score|String|
|key_value_pairs|JSON|
|last_update_time|String|
|managed|String|
|name (PK)|String|
|owner_type|String|
|score_reason|String|