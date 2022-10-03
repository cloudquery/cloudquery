\set check_id 'SecretsManager.1'
\echo "Executing check SecretsManager.1"
\ir ../queries/secretsmanager/secrets_should_have_automatic_rotation_enabled.sql

\set check_id 'SecretsManager.2'
\echo "Executing check SecretsManager.2"
\ir ../queries/secretsmanager/secrets_configured_with_automatic_rotation_should_rotate_successfully.sql

\set check_id 'SecretsManager.3'
\echo "Executing check SecretsManager.3"
\ir ../queries/secretsmanager/remove_unused_secrets_manager_secrets.sql

\set check_id 'SecretsManager.4'
\echo "Executing check SecretsManager.4"
\ir ../queries/secretsmanager/secrets_should_be_rotated_within_a_specified_number_of_days.sql
