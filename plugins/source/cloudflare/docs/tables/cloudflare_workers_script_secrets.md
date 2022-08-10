
# Table: cloudflare_workers_script_secrets
WorkersSecret contains the name and type of the secret.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|workers_script_cq_id|uuid|Unique CloudQuery ID of cloudflare_workers_scripts table (FK)|
|account_id|text|The Account ID of the resource.|
|name|text|Secret name|
|type|text|Secret type|
