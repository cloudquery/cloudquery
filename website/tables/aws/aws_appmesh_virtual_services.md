# Table: aws_appmesh_virtual_services

This table shows data for AWS App Mesh Virtual Services.

https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualServiceData.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_appmesh_meshes](aws_appmesh_meshes).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|mesh_name|`utf8`|
|metadata|`json`|
|spec|`json`|
|status|`json`|
|virtual_service_name|`utf8`|