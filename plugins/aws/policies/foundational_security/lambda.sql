\set check_id 'Lambda.1'
\echo "Executing check Lambda.1"
\i queries/lambda/lambda_function_prohibit_public_access.sql

\set check_id 'Lambda.2'
\echo "Executing check Lambda.2"
\i queries/lambda/lambda_functions_should_use_supported_runtimes.sql
