# Table: k8s_certificates_signing_requests

This table shows data for Kubernetes (K8s) Certificates Signing Requests.

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
|spec_request|`binary`|
|spec_signer_name|`utf8`|
|spec_expiration_seconds|`int64`|
|spec_usages|`list<item: utf8, nullable>`|
|spec_username|`utf8`|
|spec_uid|`utf8`|
|spec_groups|`list<item: utf8, nullable>`|
|spec_extra|`json`|
|status_conditions|`json`|
|status_certificate|`binary`|