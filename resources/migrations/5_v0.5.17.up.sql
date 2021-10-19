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



