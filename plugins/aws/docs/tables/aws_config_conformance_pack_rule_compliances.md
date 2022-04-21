
# Table: aws_config_conformance_pack_rule_compliances
Compliance information of one or more AWS Config rules within a conformance pack
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|conformance_pack_cq_id|uuid|Unique CloudQuery ID of aws_config_conformance_packs table (FK)|
|compliance_type|text|Compliance of the AWS Config rule|
|config_rule_name|text|Name of the config rule.|
|controls|text[]|Controls for the conformance pack|
|config_rule_invoked_time|timestamp without time zone|The time when AWS Config rule evaluated AWS resource.|
|resource_id|text|The ID of the evaluated AWS resource.|
|resource_type|text|The type of AWS resource that was evaluated.|
|ordering_timestamp|timestamp without time zone|The time of the event that triggered the evaluation of your AWS resources.|
|result_recorded_time|timestamp without time zone|The time when AWS Config recorded the evaluation result.|
|annotation|text|Supplementary information about how the evaluation determined the compliance.|
