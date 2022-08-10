
# Table: gcp_sql_instance_settings_deny_maintenance_periods
Deny Maintenance Periods This specifies a date range during when all CSA rollout will be denied
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique ID of gcp_sql_instances table (FK)|
|instance_name|text||
|end_date|text|"deny maintenance period" end date If the year of the end date is empty, the year of the start date also must be empty In this case, it means the deny maintenance period recurs every year The date is in format yyyy-mm-dd ie, 2020-11-01, or mm-dd, ie|
|start_date|text|"deny maintenance period" start date If the year of the start date is empty, the year of the end date also must be empty In this case, it means the deny maintenance period recurs every year The date is in format yyyy-mm-dd ie, 2020-11-01, or mm-dd, ie|
|time|text|Time in UTC when the "deny maintenance period" starts on start_date and ends on end_date The time is in format: HH:mm:SS, ie|
