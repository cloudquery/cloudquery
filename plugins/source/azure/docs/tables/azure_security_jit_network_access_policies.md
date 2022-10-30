# Table: azure_security_jit_network_access_policies



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
|name|String|
|type|String|
|kind|String|
|location|String|
|virtual_machines|JSON|
|requests|JSON|
|provisioning_state|String|