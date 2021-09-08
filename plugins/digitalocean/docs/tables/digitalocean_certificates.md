
# Table: digitalocean_certificates
Certificate represents a DigitalOcean certificate configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|A unique ID that can be used to identify and reference a certificate.|
|name|text|A unique human-readable name referring to a certificate.|
|dns_names|text[]|An array of fully qualified domain names (FQDNs) for which the certificate was issued.|
|not_after|text|A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.|
|s_h_a1_fingerprint|text|A unique identifier generated from the SHA-1 fingerprint of the certificate.|
|created|text|A time value given in ISO8601 combined date and time format that represents when the certificate was created.|
|state|text|A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.|
|type|text|A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.|
