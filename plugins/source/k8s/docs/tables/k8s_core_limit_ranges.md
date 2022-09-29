# Table: k8s_core_limit_ranges


The primary key for this table is **uid**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|spec_limits|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|