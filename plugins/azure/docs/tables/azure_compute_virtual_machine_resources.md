
# Table: azure_compute_virtual_machine_resources
VirtualMachineExtension describes a Virtual Machine Extension
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_cq_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_id|text|ID of azure_compute_virtual_machines table (FK)|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
