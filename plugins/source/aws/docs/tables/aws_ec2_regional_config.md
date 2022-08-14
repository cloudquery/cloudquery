
# Table: aws_ec2_regional_config
Ec2 Regional Config defines common default configuration for ec2 service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|ebs_encryption_enabled_by_default|boolean|Indicates whether EBS encryption by default is enabled for your account in the current Region.|
|ebs_default_kms_key_id|text|The Amazon Resource Name (ARN) of the default CMK for encryption by default.|
