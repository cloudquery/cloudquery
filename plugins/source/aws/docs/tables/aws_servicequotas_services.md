# Table: aws_servicequotas_services

https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html

The composite primary key for this table is (**account_id**, **region**, **service_code**, **service_name**).

## Relations

The following tables depend on aws_servicequotas_services:
  - [aws_servicequotas_quotas](aws_servicequotas_quotas.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|service_code (PK)|String|
|service_name (PK)|String|