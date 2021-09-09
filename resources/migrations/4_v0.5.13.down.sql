UPDATE "aws_access_analyzer_analyzer_finding_sources" SET "detail_access_point_arn" = '' WHERE "detail_access_point_arn" IS NULL;

ALTER TABLE IF EXISTS "aws_access_analyzer_analyzer_finding_sources" ALTER COLUMN "detail_access_point_arn" SET NOT NULL;

ALTER TABLE "aws_waf_web_acls"
    ADD COLUMN IF NOT EXISTS region text;