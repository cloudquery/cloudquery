# Table: k8s_core_events



The primary key for this table is **uid**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|uid (PK)|String|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|involved_object|JSON|
|reason|String|
|message|String|
|source|JSON|
|first_timestamp|Timestamp|
|last_timestamp|Timestamp|
|count|Int|
|type|String|
|event_time|JSON|
|series|JSON|
|action|String|
|related|JSON|
|reporting_component|String|
|reporting_instance|String|