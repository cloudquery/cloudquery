ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN IF EXISTS region text

ALTER TABLE IF EXISTS "aws_s3_bucket_cors_rules" ADD CONSTRAINT "aws_s3_bucket_cors_rules_pk" UNIQUE ("bucket_cq_id", "id")
