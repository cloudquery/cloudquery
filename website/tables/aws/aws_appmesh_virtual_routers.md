# Table: aws_appmesh_virtual_routers

This table shows data for AWS App Mesh Virtual Routers.

https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualRouterData.html

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
|virtual_router_name|`utf8`|