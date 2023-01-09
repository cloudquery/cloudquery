# Table: pagerduty_escalation_policies

https://developer.pagerduty.com/api-reference/51b21014a4f5a-list-escalation-policies

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|html_url|String|
|type|String|
|summary|String|
|self|String|
|name|String|
|escalation_rules|JSON|
|services|JSON|
|num_loops|Int|
|teams|JSON|
|description|String|
|on_call_handoff_notifications|String|