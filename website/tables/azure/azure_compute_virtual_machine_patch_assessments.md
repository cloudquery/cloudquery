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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|assessment_activity_id (PK)|String|
|available_patches|JSON|
|critical_and_security_patch_count|Int|
|error|JSON|
|other_patch_count|Int|
|reboot_pending|Bool|
|start_date_time|Timestamp|
|status|String|