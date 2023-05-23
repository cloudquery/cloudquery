# Table: gcp_dns_resource_record_sets

This table shows data for GCP DNS Resource Record Sets.

https://cloud.google.com/dns/docs/reference/v1/resourceRecordSets

The composite primary key for this table is (**project_id**, **name**, **type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|kind|String|
|name (PK)|String|
|routing_policy|JSON|
|rrdatas|StringArray|
|signature_rrdatas|StringArray|
|ttl|Int|
|type (PK)|String|