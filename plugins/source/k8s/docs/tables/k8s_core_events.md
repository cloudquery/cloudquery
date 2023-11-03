# Table: k8s_core_events

This table shows data for Kubernetes (K8s) Core Events.

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
|involved_object|`json`|
|reason|`utf8`|
|message|`utf8`|
|source|`json`|
|first_timestamp|`timestamp[us, tz=UTC]`|
|last_timestamp|`timestamp[us, tz=UTC]`|
|count|`int64`|
|type|`utf8`|
|event_time|`json`|
|series|`json`|
|action|`utf8`|
|related|`json`|
|reporting_component|`utf8`|
|reporting_instance|`utf8`|