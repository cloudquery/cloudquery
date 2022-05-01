ALTER TABLE IF EXISTS "aws_sns_topics" ADD COLUMN IF NOT EXISTS "tags" jsonb;
