\echo "Executing Information Exchange Policies and Procedures"
\set check_id "0662.09sCSPOrganizational.2 - 09.s - 1"
\echo :check_id
\i queries/web/app_client_cert_disabled.sql
\set check_id "0901.09s1Organizational.1 - 09.s - 1"
\echo :check_id
\i queries/web/cors_should_not_allow_every_resource_to_access_your_web_applications.sql
\set check_id "0902.09s2Organizational.13 - 09.s - 1"
\echo :check_id
\i queries/web/cors_should_not_allow_every_resource_to_access_your_function_apps.sql
\set check_id "0911.09s1Organizational.2 - 09.s - 1"
\echo :check_id
\i queries/web/cors_should_not_allow_every_resource_to_access_your_api_app.sql
\set check_id "0912.09s1Organizational.4 - 09.s - 1"
\echo :check_id
\i queries/web/remote_debugging_should_be_turned_off_for_web_applications.sql
\set check_id "0913.09s1Organizational.5 - 09.s - 1"
\echo :check_id
\i queries/web/remote_debugging_should_be_turned_off_for_function_apps.sql
\set check_id "0914.09s1Organizational.6 - 09.s - 1"
\echo :check_id
\i queries/web/remote_debugging_should_be_turned_off_for_api_apps.sql
