# Table: aws_detective_graph_members

This table shows data for Amazon Detective Graph Members.

https://docs.aws.amazon.com/detective/latest/APIReference/API_GetMembers.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The composite primary key for this table is (**request_region**, **account_id**, **graph_arn**).

## Relations

This table depends on [aws_detective_graphs](aws_detective_graphs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region (PK)|`utf8`|
|account_id (PK)|`utf8`|
|administrator_id|`utf8`|
|datasource_package_ingest_states|`json`|
|disabled_reason|`utf8`|
|email_address|`utf8`|
|graph_arn (PK)|`utf8`|
|invitation_type|`utf8`|
|invited_time|`timestamp[us, tz=UTC]`|
|master_id|`utf8`|
|percent_of_graph_utilization|`float64`|
|percent_of_graph_utilization_updated_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|updated_time|`timestamp[us, tz=UTC]`|
|volume_usage_by_datasource_package|`json`|
|volume_usage_in_bytes|`int64`|
|volume_usage_updated_time|`timestamp[us, tz=UTC]`|