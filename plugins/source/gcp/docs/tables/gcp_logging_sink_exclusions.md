
# Table: gcp_logging_sink_exclusions
Specifies a set of log entries that are not to be stored in Logging If your GCP resource receives a large volume of logs, you can use exclusions to reduce your chargeable logs Exclusions are processed after log sinks, so you can export log entries before they are excluded Note that organization-level and folder-level exclusions don't apply to child resources, and that you can't exclude audit log entries
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|sink_cq_id|uuid|Unique ID of gcp_logging_sinks table (FK)|
|create_time|text|The creation timestamp of the exclusionThis field may not be present for older exclusions|
|description|text|A description of this exclusion|
|disabled|boolean|If set to True, then this exclusion is disabled and it does not exclude any log entries You can update an exclusion to change the value of this field|
|filter|text|An advanced logs filter (https://cloudgooglecom/logging/docs/view/advanced-queries) that matches the log entries to be excluded By using the sample function (https://cloudgooglecom/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resourcetype=gcs_bucket severity<ERROR sample(insertId, 0|
|name|text|A client-assigned identifier, such as "load-balancer-exclusion" Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods First character has to be alphanumeric|
|update_time|text|The last update timestamp of the exclusionThis field may not be present for older exclusions|
