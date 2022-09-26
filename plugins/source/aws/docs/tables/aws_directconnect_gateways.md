# Table: aws_directconnect_gateways


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_directconnect_gateways`:
  - [`aws_directconnect_gateway_associations`](aws_directconnect_gateway_associations.md)
  - [`aws_directconnect_gateway_attachments`](aws_directconnect_gateway_attachments.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|id|String|
|amazon_side_asn|Int|
|direct_connect_gateway_name|String|
|direct_connect_gateway_state|String|
|owner_account|String|
|state_change_error|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|