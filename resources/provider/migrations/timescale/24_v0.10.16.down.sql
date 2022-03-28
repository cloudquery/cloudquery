-- Resource: ec2.images
ALTER TABLE IF EXISTS "aws_ec2_images" DROP COLUMN IF EXISTS "last_launched_time";
