# Table: aws_servicequotas_services

This table shows data for Servicequotas Services.

https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **service_code**, **service_name**).
## Relations

The following tables depend on aws_servicequotas_services:
  - [aws_servicequotas_quotas](aws_servicequotas_quotas.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|service_code|`utf8`|
|service_name|`utf8`|