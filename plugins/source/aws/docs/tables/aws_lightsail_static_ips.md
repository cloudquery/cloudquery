# Table: aws_lightsail_static_ips

This table shows data for Lightsail Static IPs.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_StaticIp.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|attached_to|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|ip_address|`utf8`|
|is_attached|`bool`|
|location|`json`|
|name|`utf8`|
|resource_type|`utf8`|
|support_code|`utf8`|