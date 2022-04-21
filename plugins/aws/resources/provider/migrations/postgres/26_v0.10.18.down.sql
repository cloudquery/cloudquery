-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "ephemeral_storage_size";
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "runtime_platform_cpu_architecture";
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "runtime_platform_os_family";

-- Resource: kms.keys
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "aws_account_id";
ALTER TABLE "aws_kms_keys" RENAME COLUMN "key_spec" TO "customer_master_key_spec";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "mac_algorithms";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "multi_region";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "multi_region_key_type";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "primary_key_arn";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "primary_key_region";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "replica_keys";
ALTER TABLE IF EXISTS "aws_kms_keys" DROP COLUMN IF EXISTS "pending_deletion_window_in_days";
