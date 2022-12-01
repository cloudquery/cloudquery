# Table: k8s_certificates_signing_requests



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
|spec_request|IntArray|
|spec_signer_name|String|
|spec_expiration_seconds|Int|
|spec_usages|StringArray|
|spec_username|String|
|spec_groups|StringArray|
|spec_extra|JSON|
|status_conditions|JSON|
|status_certificate|IntArray|