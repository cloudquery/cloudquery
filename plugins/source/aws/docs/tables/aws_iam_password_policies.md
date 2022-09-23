# Table: aws_iam_password_policies


The primary key for this table is **account_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|allow_users_to_change_password|Bool|
|expire_passwords|Bool|
|hard_expiry|Bool|
|max_password_age|Int|
|minimum_password_length|Int|
|password_reuse_prevention|Int|
|require_lowercase_characters|Bool|
|require_numbers|Bool|
|require_symbols|Bool|
|require_uppercase_characters|Bool|
|policy_exists|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|