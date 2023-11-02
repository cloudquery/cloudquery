# Table: aws_appmesh_virtual_routers

This table shows data for AWS App Mesh Virtual Routers.

https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualRouterData.html

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**, **mesh_arn**).

## Relations

This table depends on [aws_appmesh_meshes](aws_appmesh_meshes.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|mesh_arn (PK)|`utf8`|
|mesh_name|`utf8`|
|metadata|`json`|
|spec|`json`|
|status|`json`|
|virtual_router_name|`utf8`|