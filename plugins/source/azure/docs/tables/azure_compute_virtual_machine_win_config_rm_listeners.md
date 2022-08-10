
# Table: azure_compute_virtual_machine_win_config_rm_listeners
WinRMListener describes Protocol and thumbprint of Windows Remote Management listener
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_id|text|ID of azure_compute_virtual_machines table (FK)|
|protocol|text|Specifies the protocol of WinRM listener|
|certificate_url|text|This is the URL of a certificate that has been uploaded to Key Vault as a secret|
