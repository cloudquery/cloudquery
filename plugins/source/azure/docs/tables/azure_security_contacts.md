# Table: azure_security_contacts

This table shows data for Azure Security Contacts.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity@v0.9.0#Contact

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that "Notify about alerts with the following severity" is set to "High" (Automated)

```sql
SELECT
  'Ensure that "Notify about alerts with the following severity" is set to "High" (Automated)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN email IS NOT NULL AND email != '' AND alert_notifications = 'On'
  THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_contacts;
```

### Ensure "Additional email addresses" is configured with a security contact email (Automated)

```sql
SELECT
  'Ensure "Additional email addresses" is configured with a security contact email (Automated)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN email IS NOT NULL AND email != '' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_contacts;
```


