
# Table: gcp_security_secrets
Secret: A Secret is a logical secret whose value and versions can be accessed
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|id|text|The id of the secret|
|resource_name|text|The resource name of the Secret in the format `projects/*/secrets/*`|
|topics|text[]|A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions. In the format `projects/*/topics/*`|
|is_automatically_replicated|boolean|If true, the secret is automatically replicated by GCP. Otherwise, replications are user-managed.|
|create_time|text|The time at which the Secret was created|
|etag|text|Etag of the currently stored Secret|
|expire_time|text|Timestamp in UTC when the Secret is scheduled to expire.|
|labels|jsonb|The labels assigned to this Secret|
|automatic_replication_customer_managed_encryption_kms_key_name|text|If the secret is automatically replicated, contains the customer-managed-encryption kms-key-name. Only valid if 'is_automatically_replicated' is true. If null, then the secret is encrypted with a google-managed key.|
|next_rotation_time|text|Timestamp in UTC at which the Secret is scheduled to rotate|
