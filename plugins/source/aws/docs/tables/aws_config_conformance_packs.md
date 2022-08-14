
# Table: aws_config_conformance_packs
Returns details of a conformance pack.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|Amazon Resource Name (ARN) of the conformance pack.|
|conformance_pack_id|text|ID of the conformance pack.|
|conformance_pack_name|text|Name of the conformance pack.|
|conformance_pack_input_parameters|jsonb|A list of ConformancePackInputParameter objects.|
|created_by|text|AWS service that created the conformance pack.|
|delivery_s3_bucket|text|Amazon S3 bucket where AWS Config stores conformance pack templates.|
|delivery_s3_key_prefix|text|The prefix for the Amazon S3 bucket.|
|last_update_requested_time|timestamp without time zone|Last time when conformation pack update was requested.|
