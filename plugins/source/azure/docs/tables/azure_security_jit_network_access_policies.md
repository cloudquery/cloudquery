# Table: azure_security_jit_network_access_policies

This table shows data for Azure Security Jit Network Access Policies.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/jit-network-access-policies/list?tabs=HTTP#jitnetworkaccesspolicy

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|kind|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Management ports of virtual machines should be protected with just-in-time network access control

```sql
WITH
  jit_vms
    AS (
      SELECT
        properties->'virtualMachines'->>'id' AS vm_id
      FROM
        azure_security_jit_network_access_policies
      WHERE
        properties->>'provisioningState' = 'Succeeded'
    )
INSERT
INTO
  azure_policy_results
    (
      execution_time,
      framework,
      check_id,
      title,
      subscription_id,
      resource_id,
      status
    )
SELECT
  'Management ports of virtual machines should be protected with just-in-time network access control'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE WHEN j.vm_id = NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  azure_compute_virtual_machines AS vm LEFT JOIN jit_vms AS j ON vm.id = j.vm_id;
```


