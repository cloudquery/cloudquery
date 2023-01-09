# Table: aws_config_conformance_packs

https://docs.aws.amazon.com/config/latest/APIReference/API_ConformancePackDetail.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_config_conformance_packs:
  - [aws_config_conformance_pack_rule_compliances](aws_config_conformance_pack_rule_compliances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|conformance_pack_arn|String|
|conformance_pack_id|String|
|conformance_pack_name|String|
|conformance_pack_input_parameters|JSON|
|created_by|String|
|delivery_s3_bucket|String|
|delivery_s3_key_prefix|String|
|last_update_requested_time|Timestamp|
|template_ssm_document_details|JSON|