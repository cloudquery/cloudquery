
# Table: gcp_compute_instance_group_instances

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_group_cq_id|uuid|Unique CloudQuery ID of gcp_compute_instance_groups table (FK)|
|instance|text|The URL of the instance|
|named_ports|jsonb|The named ports that belong to this instance group|
|status|text|Status of the instance One of DEPROVISIONING, PROVISIONING, REPAIRING, RUNNING, STAGING, STOPPED, STOPPING, SUSPENDED, SUSPENDING, TERMINATED|
