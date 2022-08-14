
# Table: digitalocean_keys

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|bigint|A unique identification number for this key. Can be used to embed specific SSH key into a Droplet.|
|name|text|A human-readable display name for this key, used to easily identify the SSH keys when they are displayed.|
|fingerprint|text|A unique identifier that differentiates this key from other keys using a format that SSH recognizes. The fingerprint is created when the key is added to your account.|
|public_key|text|The entire public key string that was uploaded. Embedded into the root user's `authorized_keys` file if you include this key during Droplet creation.|
