
# Table: aws_organizations_accounts
Contains information about an AWS account that is a member of an organization
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|tags|jsonb|The AWS tags of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the account|
|email|text|The email address associated with the AWS account|
|id|text|The unique identifier (ID) of the account|
|joined_method|text|The method by which the account joined the organization|
|joined_timestamp|timestamp without time zone|The date the account became a part of the organization|
|name|text|The friendly name of the account|
|status|text|The status of the account in the organization|
