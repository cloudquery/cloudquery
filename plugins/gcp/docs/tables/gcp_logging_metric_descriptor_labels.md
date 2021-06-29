
# Table: gcp_logging_metric_descriptor_labels
A description of a label
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|metric_id|uuid|Unique ID of gcp_logging_metrics table (FK)|
|description|text|A human-readable description for the label|
|key|text|The label key|
|value_type|text|The type of data that can be assigned to the label  Possible values:   "STRING" - A variable-length string This is the default   "BOOL" - Boolean; true or false   "INT64" - A 64-bit signed integer|
