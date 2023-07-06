# Table: aws_elasticsearch_vpc_endpoints

This table shows data for Elasticsearch VPC Endpoints.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_VpcEndpoint.html

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id (PK)|`utf8`|
|domain_arn|`utf8`|
|endpoint|`utf8`|
|status|`utf8`|
|vpc_endpoint_id|`utf8`|
|vpc_endpoint_owner|`utf8`|
|vpc_options|`json`|