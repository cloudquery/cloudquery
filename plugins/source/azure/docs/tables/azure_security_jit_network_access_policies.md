# Table: azure_security_jit_network_access_policies


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|id (PK)|String|
|name|String|
|type|String|
|kind|String|
|location|String|
|virtual_machines|JSON|
|requests|JSON|
|provisioning_state|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|