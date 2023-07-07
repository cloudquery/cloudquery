# Table: aws_config_conformance_packs

This table shows data for Config Conformance Packs.

https://docs.aws.amazon.com/config/latest/APIReference/API_ConformancePackDetail.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_config_conformance_packs:
  - [aws_config_conformance_pack_rule_compliances](aws_config_conformance_pack_rule_compliances)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|conformance_pack_arn|`utf8`|
|conformance_pack_id|`utf8`|
|conformance_pack_name|`utf8`|
|conformance_pack_input_parameters|`json`|
|created_by|`utf8`|
|delivery_s3_bucket|`utf8`|
|delivery_s3_key_prefix|`utf8`|
|last_update_requested_time|`timestamp[us, tz=UTC]`|
|template_ssm_document_details|`json`|