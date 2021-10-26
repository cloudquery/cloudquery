--aws_apigatewayv2_vpc_links
ALTER TABLE IF EXISTS "aws_apigatewayv2_vpc_links"
    ADD COLUMN "vpc_link_id" TEXT;

UPDATE "aws_apigatewayv2_vpc_links"
SET "vpc_link_id" = "id";

--aws_cloudtrail_trails
ALTER TABLE IF EXISTS "aws_cloudtrail_trails"
DROP
COLUMN "tags";
ALTER TABLE IF EXISTS "aws_cloudtrail_trails"
    ADD COLUMN "home_region" TEXT;
UPDATE "aws_cloudtrail_trails"
SET "home_region" = "region";

--aws_elasticbeanstalk_environments
ALTER TABLE IF EXISTS "aws_elasticbeanstalk_environments"
DROP
COLUMN "tags";
ALTER TABLE IF EXISTS "aws_elasticbeanstalk_environments" RENAME COLUMN "name" TO "environment_name";

--aws_elasticsearch_domains
ALTER TABLE IF EXISTS "aws_elasticsearch_domains"
DROP
COLUMN "tags";

--aws_elbv2_load_balancers
ALTER TABLE IF EXISTS "aws_elbv2_load_balancers"
DROP
COLUMN "tags";

--aws_elbv2_target_groups
ALTER TABLE IF EXISTS "aws_elbv2_target_groups"
DROP
COLUMN "tags";

--aws_kms_keys
ALTER TABLE IF EXISTS "aws_kms_keys"
DROP
COLUMN "tags";
ALTER TABLE IF EXISTS "aws_kms_keys" RENAME COLUMN "id" TO "key_id";

--aws_elbv2_load_balancer_availability_zone_addresses
ALTER TABLE IF EXISTS "aws_elbv2_load_balancer_availability_zone_addresses" DROP CONSTRAINT IF EXISTS "aws_elbv2_load_balancer_availability_zone_addresses_pk";
ALTER TABLE IF EXISTS "aws_elbv2_load_balancer_availability_zone_addresses" ADD CONSTRAINT "aws_elbv2_load_balancer_availability_zone_addresses_pk" UNIQUE ("load_balancer_availability_zone_cq_id", "ip_address");

--aws_apigateway_domain_name_base_path_mappings
ALTER TABLE IF EXISTS "aws_apigateway_domain_name_base_path_mappings" DROP CONSTRAINT IF EXISTS "aws_apigateway_domain_name_base_path_mappings_pk";
ALTER TABLE IF EXISTS "aws_apigateway_domain_name_base_path_mappings" ADD CONSTRAINT "aws_apigateway_domain_name_base_path_mappings_pk" UNIQUE ("domain_name_cq_id", "rest_api_id");
