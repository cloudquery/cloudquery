# Table: k8s_storage_volume_attachments

This table shows data for Kubernetes (K8s) Storage Volume Attachments.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|kind|`utf8`|
|api_version|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|uid (PK)|`utf8`|
|resource_version|`utf8`|
|generation|`int64`|
|deletion_grace_period_seconds|`int64`|
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec_attacher|`utf8`|
|spec_source|`json`|
|spec_node_name|`utf8`|
|status_attached|`bool`|
|status_attachment_metadata|`json`|
|status_attach_error|`json`|
|status_detach_error|`json`|