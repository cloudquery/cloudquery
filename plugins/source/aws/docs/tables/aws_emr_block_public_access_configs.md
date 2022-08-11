
# Table: aws_emr_block_public_access_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|block_public_security_group_rules|boolean|Indicates whether Amazon EMR block public access is enabled or disabled.|
|classification|text|The classification within a configuration.|
|configurations|jsonb|A list of additional configurations to apply within a configuration object.|
|properties|jsonb|A set of properties specified within a configuration classification.|
|created_by_arn|text|The Amazon Resource Name that created or last modified the configuration.|
|creation_date_time|timestamp without time zone|The date and time that the configuration was created.|
