
# Table: aws_route53_hosted_zone_vpc_association_authorizations
(Private hosted zones only) A complex type that contains information about an Amazon VPC.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_cq_id|uuid|Unique CloudQuery ID of aws_route53_hosted_zones table (FK)|
|vpc_id|text|(Private hosted zones only) The ID of an Amazon VPC.|
|vpc_region|text|(Private hosted zones only) The region that an Amazon VPC was created in.|
