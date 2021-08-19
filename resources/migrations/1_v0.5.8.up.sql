ALTER TABLE "aws_cloudfront_distributions"
    DROP COLUMN IF EXISTS region;

ALTER TABLE IF EXISTS "aws_s3_bucket_cors_rules" DROP CONSTRAINT IF EXISTS "aws_s3_bucket_cors_rules_pk";

