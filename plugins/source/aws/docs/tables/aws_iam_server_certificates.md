# Table: aws_iam_server_certificates

This table shows data for IAM Server Certificates.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|arn|`utf8`|
|path|`utf8`|
|server_certificate_id|`utf8`|
|server_certificate_name|`utf8`|
|expiration|`timestamp[us, tz=UTC]`|
|upload_date|`timestamp[us, tz=UTC]`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that all the expired SSL/TLS certificates stored in AWS IAM are removed (Automated)

```sql
SELECT
  'Ensure that all the expired SSL/TLS certificates stored in AWS IAM are removed (Automated)'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN expiration < timezone('UTC', now()) THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_server_certificates;
```


