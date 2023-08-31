# Table: aws_account_alternate_contacts

This table shows data for Account Alternate Contacts.

https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html

The composite primary key for this table is (**account_id**, **alternate_contact_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|alternate_contact_type (PK)|`utf8`|
|email_address|`utf8`|
|name|`utf8`|
|phone_number|`utf8`|
|title|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Security contact information should be provided for an AWS account

```sql
SELECT
  'Security contact information should be provided for an AWS account' AS title,
  aws_iam_accounts.account_id,
  CASE WHEN alternate_contact_type IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_iam_accounts
  LEFT JOIN (
      SELECT
        *
      FROM
        aws_account_alternate_contacts
      WHERE
        alternate_contact_type = 'SECURITY'
    )
      AS account_security_contacts ON
      aws_iam_accounts.account_id = account_security_contacts.account_id;
```


