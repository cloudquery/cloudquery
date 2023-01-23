# Table: gcp_appengine_firewall_ingress_rules

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules#FirewallRule

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|priority|Int|
|action|String|
|source_range|String|
|description|String|