ALTER TABLE IF EXISTS "aws_access_analyzer_analyzer_finding_sources" ALTER COLUMN "detail_access_point_arn" DROP NOT NULL;

ALTER TABLE "aws_waf_web_acls"
    DROP COLUMN IF EXISTS region;