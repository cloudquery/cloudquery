
# Table: gcp_kms_keyrings
A KeyRing is a toplevel logical grouping of CryptoKeys.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|location|text|Location of the resource|
|create_time|timestamp without time zone|The time at which this KeyRing was created|
|name|text|The resource name for the KeyRing in the format `projects/*/locations/*/keyRings/*`|
