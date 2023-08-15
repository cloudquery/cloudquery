# Table: aws_servicequotas_services

This table shows data for Servicequotas Services.

https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html

The composite primary key for this table is (**account_id**, **region**, **service_code**, **service_name**).

## Relations

The following tables depend on aws_servicequotas_services:
  - [aws_servicequotas_quotas](aws_servicequotas_quotas)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|service_code (PK)|`utf8`|
|service_name (PK)|`utf8`|