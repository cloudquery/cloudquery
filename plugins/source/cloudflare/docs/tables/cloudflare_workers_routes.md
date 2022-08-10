
# Table: cloudflare_workers_routes
WorkerRoute is used to map traffic matching a URL pattern to a workers  API reference: https://api.cloudflare.com/#worker-routes-properties
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|zone_id|text|The Zone ID of the resource.|
|id|text|API item identifier tag|
|pattern|text|The pattern of the route.|
|enabled|boolean|Whether the route is enabled|
|script|text|Name of the script to apply when the route is matched. The route is skipped when this is blank/missing.|
