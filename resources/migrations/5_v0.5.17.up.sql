--aws_apigatewayv2_vpc_links
ALTER TABLE IF EXISTS "aws_apigatewayv2_vpc_links"
DROP
COLUMN "vpc_link_id";

--aws_cloudtrail_trails
ALTER TABLE IF EXISTS "aws_cloudtrail_trails"
    ADD COLUMN "tags" json;

-- home_region duplicates region column so we removed it
ALTER TABLE IF EXISTS "aws_cloudtrail_trails"
DROP
COLUMN "home_region";

--aws_elasticbeanstalk_environments
ALTER TABLE IF EXISTS "aws_elasticbeanstalk_environments"
    ADD COLUMN "tags" json;
ALTER TABLE IF EXISTS "aws_elasticbeanstalk_environments" RENAME COLUMN "environment_name" TO "name";

--aws_elasticsearch_domains
ALTER TABLE IF EXISTS "aws_elasticsearch_domains"
    ADD COLUMN "tags" json;

--aws_elbv2_load_balancers
ALTER TABLE IF EXISTS "aws_elbv2_load_balancers"
    ADD COLUMN "tags" json;

--aws_elbv2_target_groups
ALTER TABLE IF EXISTS "aws_elbv2_target_groups"
    ADD COLUMN "tags" json;

--aws_kms_keys
ALTER TABLE IF EXISTS "aws_kms_keys"
    ADD COLUMN "tags" json;
ALTER TABLE IF EXISTS "aws_kms_keys" RENAME COLUMN "key_id" TO "id";

--aws_elbv2_load_balancer_availability_zone_addresses
ALTER TABLE IF EXISTS "aws_elbv2_load_balancer_availability_zone_addresses" DROP CONSTRAINT IF EXISTS "aws_elbv2_load_balancer_availability_zone_addresses_pk";
ALTER TABLE IF EXISTS "aws_elbv2_load_balancer_availability_zone_addresses" ADD CONSTRAINT "aws_elbv2_load_balancer_availability_zone_addresses_pk" UNIQUE ("load_balancer_availability_zone_cq_id", "cq_id");

--aws_apigateway_domain_name_base_path_mappings
ALTER TABLE IF EXISTS "aws_apigateway_domain_name_base_path_mappings" DROP CONSTRAINT IF EXISTS "aws_apigateway_domain_name_base_path_mappings_pk";
ALTER TABLE IF EXISTS "aws_apigateway_domain_name_base_path_mappings" ADD CONSTRAINT "aws_apigateway_domain_name_base_path_mappings_pk" UNIQUE ("domain_name_cq_id", "cq_id");

--aws_emr_clusters_vpc_id
ALTER TABLE IF EXISTS "aws_emr_clusters"
ADD COLUMN "vpc_id" text;


