
# Table: azure_compute_virtual_machine_win_config_rm_listeners
WinRMListener describes Protocol and thumbprint of Windows Remote Management listener
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_cq_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_id|text|ID of azure_compute_virtual_machines table (FK)|
|protocol|text|Specifies the protocol of WinRM listener <br><br> Possible values are: <br>**http** <br><br> **https** Possible values include: 'HTTP', 'HTTPS'|
|certificate_url|text|This is the URL of a certificate that has been uploaded to Key Vault as a secret For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docsmicrosoftcom/azure/key-vault/key-vault-get-started/#add) In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  "data":"<Base64-encoded-certificate>",<br>  "dataType":"pfx",<br>  "password":"<pfx-file-password>"<br>}|
