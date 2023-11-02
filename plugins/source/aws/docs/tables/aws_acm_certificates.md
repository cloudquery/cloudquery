# Table: aws_acm_certificates

This table shows data for Amazon Certificate Manager (ACM) Certificates.

https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|certificate_authority_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|domain_name|`utf8`|
|domain_validation_options|`json`|
|extended_key_usages|`json`|
|failure_reason|`utf8`|
|imported_at|`timestamp[us, tz=UTC]`|
|in_use_by|`list<item: utf8, nullable>`|
|issued_at|`timestamp[us, tz=UTC]`|
|issuer|`utf8`|
|key_algorithm|`utf8`|
|key_usages|`json`|
|not_after|`timestamp[us, tz=UTC]`|
|not_before|`timestamp[us, tz=UTC]`|
|options|`json`|
|renewal_eligibility|`utf8`|
|renewal_summary|`json`|
|revocation_reason|`utf8`|
|revoked_at|`timestamp[us, tz=UTC]`|
|serial|`utf8`|
|signature_algorithm|`utf8`|
|status|`utf8`|
|subject|`utf8`|
|subject_alternative_names|`list<item: utf8, nullable>`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Certificate has less than 30 days to be renewed

```sql
SELECT
  'certificate has less than 30 days to be renewed' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN not_after < timezone('UTC', now()) + '30'::INTERVAL DAY THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_acm_certificates;
```

### Unused ACM certificate

```sql
SELECT
  'Unused ACM certificate' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_acm_certificates
WHERE
  array_length(in_use_by, 1) = 0;
```

### Classic Load Balancers with SSL/HTTPS listeners should use a certificate provided by AWS Certificate Manager

```sql
SELECT
  'Classic Load Balancers with SSL/HTTPS listeners should use a certificate provided by AWS Certificate Manager'
    AS title,
  lb.account_id,
  lb.arn AS resource_id,
  CASE
  WHEN li->'Listener'->>'Protocol' = 'HTTPS'
  AND aws_acm_certificates.arn IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv1_load_balancers AS lb,
  jsonb_array_elements(lb.listener_descriptions) AS li
  LEFT JOIN aws_acm_certificates ON
      aws_acm_certificates.arn = li->'Listener'->>'SSLCertificateId';
```


