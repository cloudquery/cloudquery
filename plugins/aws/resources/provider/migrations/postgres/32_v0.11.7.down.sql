-- Resource: aws.regions
ALTER TABLE IF EXISTS "aws_regions" DROP COLUMN IF EXISTS "partition";
