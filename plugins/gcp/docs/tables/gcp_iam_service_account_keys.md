
# Table: gcp_iam_service_account_keys
Represents a service account key A service account has two sets of key-pairs: user-managed, and system-managed User-managed key-pairs can be created and deleted by users Users are responsible for rotating these keys periodically to ensure security of their service accounts Users retain the private key of these key-pairs, and Google retains ONLY the public key System-managed keys are automatically rotated by Google, and are used for signing for a maximum of two weeks The rotation process is probabilistic, and usage of the new key will gradually ramp up and down over the key's lifetime If you cache the public key set for a service account, we recommend that you update the cache every 15 minutes User-managed keys can be added and removed at any time, so it is important to update the cache frequently For Google-managed keys, Google will publish a key at least 6 hours before it is first used for signing and will keep publishing it for at least 6 hours after it was last used for signing Public keys for all service accounts are also published at the OAuth2 Service Account API
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_account_id|uuid|Unique ID of gcp_iam_service_accounts table (FK)|
|key_algorithm|text|Specifies the algorithm (and possibly key size) for the key  Possible values:   "KEY_ALG_UNSPECIFIED" - An unspecified key algorithm   "KEY_ALG_RSA_1024" - 1k RSA Key   "KEY_ALG_RSA_2048" - 2k RSA Key|
|key_origin|text|The key origin  Possible values:   "ORIGIN_UNSPECIFIED" - Unspecified key origin   "USER_PROVIDED" - Key is provided by user   "GOOGLE_PROVIDED" - Key is provided by Google|
|key_type|text|The key type  Possible values:   "KEY_TYPE_UNSPECIFIED" - Unspecified key type The presence of this in the message will immediately result in an error   "USER_MANAGED" - User-managed keys (managed and rotated by the user)   "SYSTEM_MANAGED" - System-managed keys (managed and rotated by Google)|
|name|text|The resource name of the service account key in the following format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}`|
|valid_after_time|timestamp without time zone|The key can be used after this timestamp|
|valid_before_time|timestamp without time zone|The key can be used before this timestamp For system-managed key pairs, this timestamp is the end time for the private key signing operation The public key could still be used for verification for a few hours after this time|
