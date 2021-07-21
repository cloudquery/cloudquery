
# Table: azure_compute_virtual_machine_secret_vault_certificates
VaultCertificate describes a single certificate reference in a Key Vault, and where the certificate should reside on the VM
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_secret_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machine_secrets table (FK)|
|certificate_url|text|This is the URL of a certificate that has been uploaded to Key Vault as a secret For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docsmicrosoftcom/azure/key-vault/key-vault-get-started/#add) In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  "data":"<Base64-encoded-certificate>",<br>  "dataType":"pfx",<br>  "password":"<pfx-file-password>"<br>}|
|certificate_store|text|UppercaseThumbprint&gt;crt for the X509 certificate file and &lt;UppercaseThumbprint&gt;prv for private key Both of these files are pem formatted|
