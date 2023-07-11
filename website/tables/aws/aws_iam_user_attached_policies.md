# Table: aws_iam_user_attached_policies

This table shows data for IAM User Attached Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The composite primary key for this table is (**account_id**, **user_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|user_arn (PK)|`utf8`|
|policy_name (PK)|`utf8`|
|user_id|`utf8`|
|policy_arn|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### IAM users should not have IAM policies attached

```sql
SELECT
  DISTINCT
  'IAM users should not have IAM policies attached' AS title,
  aws_iam_users.account_id,
  arn AS resource_id,
  CASE
  WHEN aws_iam_user_attached_policies.user_arn IS NOT NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_users
  LEFT JOIN aws_iam_user_attached_policies ON
      aws_iam_users.arn = aws_iam_user_attached_policies.user_arn;
```


