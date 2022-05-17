-- Resource: aws.regions
ALTER TABLE IF EXISTS "aws_regions" ADD COLUMN IF NOT EXISTS "partition" text;
