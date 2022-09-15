
# Table: cloudflare_workers_script_cron_triggers
WorkerCronTrigger holds an individual cron schedule for a worker.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|workers_script_cq_id|uuid|Unique CloudQuery ID of cloudflare_workers_scripts table (FK)|
|account_id|text|The Account ID of the resource.|
|cron|text|Raw cron expression|
|created_on|timestamp without time zone|When the Cron was created|
|modified_on|timestamp without time zone|When the Cron was last modified|
