# Table: azure_sql_servers

This table shows data for Azure SQL Servers.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_servers:
  - [azure_sql_server_admins](azure_sql_server_admins)
  - [azure_sql_server_advanced_threat_protection_settings](azure_sql_server_advanced_threat_protection_settings)
  - [azure_sql_server_blob_auditing_policies](azure_sql_server_blob_auditing_policies)
  - [azure_sql_server_databases](azure_sql_server_databases)
  - [azure_sql_server_encryption_protectors](azure_sql_server_encryption_protectors)
  - [azure_sql_server_firewall_rules](azure_sql_server_firewall_rules)
  - [azure_sql_server_security_alert_policies](azure_sql_server_security_alert_policies)
  - [azure_sql_server_virtual_network_rules](azure_sql_server_virtual_network_rules)
  - [azure_sql_server_vulnerability_assessments](azure_sql_server_vulnerability_assessments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Azure Active Directory Admin is configured (Automated)

```sql
WITH
  ad_admins_count
    AS (
      SELECT
        ass._cq_id, count(*) AS admins_count
      FROM
        azure_sql_servers AS ass
        LEFT JOIN azure_sql_server_admins AS assa ON
            ass._cq_id = assa._cq_parent_id
      WHERE
        assa.properties->>'administratorType' = 'ActiveDirectory'
      GROUP BY
        ass._cq_id, assa.properties->>'administratorType'
    )
SELECT
  'Ensure that Azure Active Directory Admin is configured (Automated)' AS title,
  s.subscription_id,
  s.id,
  CASE
  WHEN a.admins_count IS NULL OR a.admins_count = 0 THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s LEFT JOIN ad_admins_count AS a ON s._cq_id = a._cq_id;
```

### Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)

```sql
SELECT
  'Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN atp.properties->>'state' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  JOIN azure_sql_server_advanced_threat_protection_settings AS atp ON
      s._cq_id = atp._cq_parent_id;
```

### Ensure that "Auditing" is set to "On" (Automated)

```sql
SELECT
  'Ensure that "Auditing" is set to "On" (Automated)' AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN assdbap.properties->>'state' != 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS assd ON s._cq_id = assd._cq_parent_id
  LEFT JOIN azure_sql_server_database_blob_auditing_policies AS assdbap ON
      assd._cq_id = assdbap._cq_parent_id;
```

### Ensure that "Auditing" Retention is "greater than 90 days" (Automated)

```sql
SELECT
  'Ensure that "Auditing" Retention is "greater than 90 days" (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN (assdbap.properties->'retentionDays')::INT8 < 90 THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_blob_auditing_policies AS assdbap ON
      s._cq_id = assdbap._cq_parent_id;
```

### Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)

```sql
SELECT
  'Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)'
    AS title,
  s.subscription_id,
  asd.id AS database_id,
  CASE
  WHEN tde.properties->>'state' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS asd ON s._cq_id = asd._cq_parent_id
  LEFT JOIN azure_sql_transparent_data_encryptions AS tde ON
      asd._cq_id = tde._cq_parent_id
WHERE
  asd.name != 'master';
```

### Long-term geo-redundant backup should be enabled for Azure SQL Databases

```sql
SELECT
  'Long-term geo-redundant backup should be enabled for Azure SQL Databases'
    AS title,
  s.subscription_id,
  rp.id,
  CASE
  WHEN rp.id IS NULL
  OR (
      rp.properties->>'weeklyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND rp.properties->>'monthlyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND rp.properties->>'yearlyRetention' IS NOT DISTINCT FROM 'PT0S'
    )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS asd ON s._cq_id = asd._cq_parent_id
  LEFT JOIN azure_sql_server_database_long_term_retention_policies AS rp ON
      asd._cq_id = rp._cq_parent_id;
```

### Vulnerability assessment should be enabled on your SQL servers

```sql
WITH
  protected_servers
    AS (
      SELECT
        s.id AS server_id
      FROM
        azure_sql_servers AS s
        LEFT JOIN azure_sql_server_vulnerability_assessments AS va ON
            s._cq_id = va._cq_parent_id
      WHERE
        (va.properties->'recurringScans'->>'isEnabled')::BOOL IS true
    )
SELECT
  'Vulnerability assessment should be enabled on your SQL servers' AS title,
  i.subscription_id,
  i.id AS instance_id,
  CASE
  WHEN p.server_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS i LEFT JOIN protected_servers AS p ON p.server_id = i.id;
```

### SQL databases should have vulnerability findings resolved

```sql
WITH
  safe_dbs
    AS (
      SELECT
        s.id AS sql_database_id
      FROM
        azure_sql_server_database_vulnerability_assessment_scans AS s
        JOIN (
            SELECT
              _cq_id, max((properties->>'endTime')::TIMESTAMP) AS max_dt
            FROM
              azure_sql_server_database_vulnerability_assessment_scans
            GROUP BY
              _cq_id
          )
            AS t ON
            s._cq_id = t._cq_id
            AND (properties->>'endTime')::TIMESTAMP = t.max_dt
      WHERE
        (s.properties->>'numberOfFailedSecurityChecks')::INT8 = 0
    )
SELECT
  'SQL databases should have vulnerability findings resolved' AS title,
  s.subscription_id,
  d.id,
  CASE
  WHEN d.id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS d ON s._cq_id = d._cq_parent_id
  LEFT JOIN safe_dbs AS sd ON d.id = sd.sql_database_id;
```

### SQL Server should use a virtual network service endpoint

```sql
WITH
  subs
    AS (
      SELECT
        subscription_id,
        jsonb_array_elements(properties->'subnets') AS subnet,
        properties->>'provisioningState' AS provisioning_state
      FROM
        azure_network_virtual_networks
    ),
  secured_servers
    AS (
      SELECT
        s._cq_id
      FROM
        azure_sql_servers AS s
        LEFT JOIN azure_sql_server_virtual_network_rules AS r ON
            s._cq_id = r._cq_parent_id
        LEFT JOIN subs ON
            r.properties->>'virtualNetworkSubnetId' = subs.subnet->>'id'
      WHERE
        (r.properties->'virtualNetworkSubnetId') IS NOT NULL
        AND subs.provisioning_state = 'Succeeded'
    )
SELECT
  'SQL Server should use a virtual network service endpoint' AS title,
  subscription_id,
  id,
  CASE
  WHEN ss._cq_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s LEFT JOIN secured_servers AS ss ON s._cq_id = ss._cq_id;
```

### Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)

```sql
SELECT
  'Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN p.kind != 'azurekeyvault'
  OR p.properties->>'serverKeyType' IS DISTINCT FROM 'AzureKeyVault'
  OR (p.properties->>'uri') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_encryption_protectors AS p ON
      s._cq_id = p._cq_parent_id;
```

### Auditing on SQL server should be enabled

```sql
SELECT
  'Auditing on SQL server should be enabled' AS title,
  sub.id,
  sub.display_name AS subscription_name,
  CASE
  WHEN azure_sql_server_blob_auditing_policies._cq_parent_id
  = azure_sql_servers._cq_id
  AND sub.id = azure_sql_servers.subscription_id
  AND azure_sql_server_blob_auditing_policies.properties->>'state' = 'Disabled'
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_server_blob_auditing_policies,
  azure_sql_servers,
  azure_subscription_subscriptions AS sub;
```

### Ensure that Vulnerability Assessment (VA) is enabled on a SQL server by setting a Storage Account (Automated)

```sql
SELECT
  'Ensure that Vulnerability Assessment (VA) is enabled on a SQL server by setting a Storage Account (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN (a.properties->>'storageContainerPath') IS NULL
  OR a.properties->>'storageContainerPath' = ''
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_vulnerability_assessments AS a ON
      s._cq_id = a._cq_parent_id;
```

### Ensure that VA setting Periodic Recurring Scans is enabled on a SQL server (Automated)

```sql
SELECT
  'Ensure that VA setting Periodic Recurring Scans is enabled on a SQL server (Automated)'
    AS title,
  s.subscription_id,
  s.id,
  CASE
  WHEN (a.properties->'recurringScans'->>'isEnabled')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_vulnerability_assessments AS a ON
      s._cq_id = a._cq_parent_id;
```

### Ensure that VA setting "Also send email notifications to admins and subscription owners" is set for a SQL server (Automated)

```sql
SELECT
  'Ensure that VA setting "Also send email notifications to admins and subscription owners" is set for a SQL server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN (a.properties->'recurringScans'->>'emailSubscriptionAdmins')::BOOL
  IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_vulnerability_assessments AS a ON
      s._cq_id = a._cq_parent_id;
```

### Ensure that VA setting Send scan reports to is configured for a SQL server (Automated)

```sql
WITH
  vulnerability_emails
    AS (
      SELECT
        id,
        unnest((v.properties->'recurringScans'->>'emails')::STRING[]) AS emails
      FROM
        azure_sql_server_vulnerability_assessments AS v
    ),
  emails_count
    AS (
      SELECT
        id, count(emails) AS emails_number
      FROM
        vulnerability_emails
      GROUP BY
        id
    )
SELECT
  'Ensure that VA setting Send scan reports to is configured for a SQL server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN c.emails_number = 0 OR c.emails_number IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_vulnerability_assessments AS sv ON
      s._cq_id = sv._cq_parent_id
  LEFT JOIN emails_count AS c ON sv.id = c.id;
```


