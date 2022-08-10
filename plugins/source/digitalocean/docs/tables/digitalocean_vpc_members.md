
# Table: digitalocean_vpc_members
Resources that are members of a VPC.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_cq_id|uuid|Unique CloudQuery ID of digitalocean_vpcs table (FK)|
|type|text|The resource type of the URN associated with the VPC..|
|id|text|A unique ID that can be used to identify the resource that is a member of the VPC.|
|urn|text|The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.|
|name|text|The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format.|
