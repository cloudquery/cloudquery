# Table: aws_config_conformance_pack_rule_compliances

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_config_conformance_packs](aws_config_conformance_packs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|compliance_type|String|
|config_rule_name|String|
|controls|StringArray|
|config_rule_invoked_time|Timestamp|
|evaluation_result_identifier|JSON|
|result_recorded_time|Timestamp|
|annotation|String|