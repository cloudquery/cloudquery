# Table: gcp_appengine_firewall_ingress_rules

This table shows data for GCP App Engine Firewall Ingress Rules.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules#FirewallRule

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|priority|int64|
|action|utf8|
|source_range|utf8|
|description|utf8|