
# Table: azure_compute_virtual_machine_secret_vault_certificates
VaultCertificate describes a single certificate reference in a Key Vault, and where the certificate should reside on the VM.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_secret_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machine_secrets table (FK)|
|certificate_url|text|This is the URL of a certificate that has been uploaded to Key Vault as a secret|
|certificate_store|text|For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate should be added|
