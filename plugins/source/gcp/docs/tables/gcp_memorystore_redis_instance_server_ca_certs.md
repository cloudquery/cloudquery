
# Table: gcp_memorystore_redis_instance_server_ca_certs
List of server CA certificates for the instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|redis_instance_cq_id|uuid|Unique CloudQuery ID of gcp_memorystore_redis_instances table (FK)|
|cert|text|PEM representation|
|create_time|text|The time when the certificate was created in RFC 3339 format|
|expire_time|text|The time when the certificate expires in RFC 3339 format|
|serial_number|text|Serial number, as extracted from the certificate|
|sha1_fingerprint|text|Sha1 Fingerprint of the certificate|
