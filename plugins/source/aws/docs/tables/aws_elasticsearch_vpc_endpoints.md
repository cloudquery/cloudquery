# Table: aws_elasticsearch_vpc_endpoints

This table shows data for Elasticsearch VPC Endpoints.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_VpcEndpoint.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|domain_arn|`utf8`|
|endpoint|`utf8`|
|status|`utf8`|
|vpc_endpoint_id|`utf8`|
|vpc_endpoint_owner|`utf8`|
|vpc_options|`json`|