# Table: azure_compute_virtual_machine_patch_assessments

This table shows data for Azure Compute Virtual Machine Patch Assessments.

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/assess-patches?tabs=HTTP#virtualmachineassesspatchesresult.

This will begin patch assessments on available virtual machines and can take long to complete.

Not available for all VMs. More at https://learn.microsoft.com/en-us/azure/virtual-machines/automatic-vm-guest-patching#requirements-for-enabling-automatic-vm-guest-patching

The primary key for this table is **assessment_activity_id**.

## Relations

This table depends on [azure_compute_virtual_machines](azure_compute_virtual_machines).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|assessment_activity_id (PK)|`utf8`|
|available_patches|`json`|
|critical_and_security_patch_count|`int64`|
|error|`json`|
|other_patch_count|`int64`|
|reboot_pending|`bool`|
|start_date_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|