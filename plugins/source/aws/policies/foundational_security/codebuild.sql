\set check_id 'CodeBuild.1'
\echo "Executing check CodeBuild.1"
\ir ../queries/codebuild/check_oauth_usage_for_sources.sql

\set check_id 'CodeBuild.2'
\echo "Executing check CodeBuild.2"
\ir ../queries/codebuild/check_environment_variables.sql
