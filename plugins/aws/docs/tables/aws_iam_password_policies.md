
# Table: aws_iam_password_policies
Contains information about the account password policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|allow_users_to_change_password|boolean|Specifies whether IAM users are allowed to change their own password. |
|expire_passwords|boolean|Indicates whether passwords in the account expire. Returns true if MaxPasswordAge contains a value greater than 0. Returns false if MaxPasswordAge is 0 or not present. |
|hard_expiry|boolean|Specifies whether IAM users are prevented from setting a new password after their password has expired. |
|max_password_age|integer|The number of days that an IAM user password is valid. |
|minimum_password_length|integer|Minimum length to require for IAM user passwords. |
|password_reuse_prevention|integer|Specifies the number of previous passwords that IAM users are prevented from reusing. |
|require_lowercase_characters|boolean|Specifies whether IAM user passwords must contain at least one lowercase character (a to z). |
|require_numbers|boolean|Specifies whether IAM user passwords must contain at least one numeric character (0 to 9). |
|require_symbols|boolean|Specifies whether IAM user passwords must contain at least one of the following symbols: ! @ # $ % ^ & * ( ) _ + - = [ ] { } | ' |
|require_uppercase_characters|boolean|Specifies whether IAM user passwords must contain at least one uppercase character (A to Z). |
|policy_exists|boolean|Specifies whether IAM user passwords configuration exists|
