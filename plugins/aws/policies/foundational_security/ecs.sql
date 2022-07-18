\set check_id 'ECS.1'
\echo "Executing check ECS.1"
\i queries/ecs/task_definitions_secure_networking.sql

\set check_id 'ECS.2'
\echo "Executing check ECS.2"
\i queries/ecs/ecs_services_with_public_ips.sql
