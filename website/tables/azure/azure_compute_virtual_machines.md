# Table: azure_compute_virtual_machines

This table shows data for Azure Compute Virtual Machines.

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/list?tabs=HTTP#virtualmachine

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machines:
  - [azure_compute_virtual_machine_extensions](azure_compute_virtual_machine_extensions)
  - [azure_compute_virtual_machine_patch_assessments](azure_compute_virtual_machine_patch_assessments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|instance_view|`json`|
|location|`utf8`|
|extended_location|`json`|
|identity|`json`|
|plan|`json`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|resources|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that ''OS and Data'' disks are encrypted with CMK (Automated)

```sql
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
  e'Ensure that \'OS and Data\' disks are encrypted with CMK (Automated)'
    AS title,
  v.subscription_id AS subscription_id,
  v.id AS resource_id,
  CASE
  WHEN d.properties->'encryption'->>'type' NOT LIKE '%CustomerKey%' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_compute_virtual_machines AS v
  JOIN azure_compute_disks AS d ON
      lower(v.id) = lower(d.properties->>'managedBy');
```

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

### Ensure Virtual Machines are utilizing Managed Disks (Manual)

```sql
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
  'Ensure Virtual Machines are utilizing Managed Disks (Manual)' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->'storageProfile'->'osDisk'->'managedDisk'->'id') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_compute_virtual_machines;
```


