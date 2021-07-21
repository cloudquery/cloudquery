
# Table: azure_compute_virtual_machine_network_interfaces
NetworkInterfaceReference describes a network interface reference
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_id|text|ID of azure_compute_virtual_machines table (FK)|
|network_interface_reference_properties_primary|boolean|Specifies the primary network interface in case the virtual machine has more than 1 network interface|
|id|text|Resource Id|
