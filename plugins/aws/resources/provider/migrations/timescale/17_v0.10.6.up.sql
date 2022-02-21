ALTER TABLE IF EXISTS "aws_s3_buckets" 
	ADD COLUMN IF NOT EXISTS "ownership_controls" text[]