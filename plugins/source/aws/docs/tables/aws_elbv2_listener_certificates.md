# Table: aws_elbv2_listener_certificates

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Listener Certificates.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Certificate.html

The composite primary key for this table is (**listener_arn**, **arn**).

## Relations

This table depends on [aws_elbv2_listeners](aws_elbv2_listeners.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|listener_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|certificate_arn|`utf8`|
|is_default|`bool`|