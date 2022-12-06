# Table: k8s_storage_volume_attachments



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
|spec_attacher|String|
|spec_source|JSON|
|spec_node_name|String|
|status_attached|Bool|
|status_attachment_metadata|JSON|
|status_attach_error|JSON|
|status_detach_error|JSON|