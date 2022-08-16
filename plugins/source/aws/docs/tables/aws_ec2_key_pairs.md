
# Table: aws_ec2_key_pairs
Describes an EC2 Key Pair.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|create_time|timestamp without time zone|The date and time when the key was created in ISO 8601 date-time format.|
|key_fingerprint|text|The fingerprint of the private key digest.|
|key_name|text|The name of the key pair.|
|key_pair_id|text|The ID of the key pair.|
|tags|jsonb|Any tags assigned to the key pair.|
