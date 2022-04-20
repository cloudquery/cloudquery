-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "ephemeral_storage_size";
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "runtime_platform_cpu_architecture";
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "runtime_platform_os_family";
