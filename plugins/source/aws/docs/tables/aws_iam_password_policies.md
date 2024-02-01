# Table: aws_iam_password_policies

This table shows data for IAM Password Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_PasswordPolicy.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|allow_users_to_change_password|`bool`|
|expire_passwords|`bool`|
|hard_expiry|`bool`|
|max_password_age|`int64`|
|minimum_password_length|`int64`|
|password_reuse_prevention|`int64`|
|require_lowercase_characters|`bool`|
|require_numbers|`bool`|
|require_symbols|`bool`|
|require_uppercase_characters|`bool`|
|policy_exists|`bool`|