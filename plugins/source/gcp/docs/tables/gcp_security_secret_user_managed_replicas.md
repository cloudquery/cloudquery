
# Table: gcp_security_secret_user_managed_replicas
Describes user-managed replicas of this secret. Empty for automatically replicated secrets
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|secret_cq_id|uuid|Unique CloudQuery ID of gcp_security_secrets table (FK)|
|customer_managed_encryption_kms_key_name|text|If the replica is encrypted with customer-managed encryption, contains the kms key name. If the column is NULL, the replica is encrypted with a google-managed key|
|location|text|The canonical IDs of the location to replicate data. For example: "us-east1"|
