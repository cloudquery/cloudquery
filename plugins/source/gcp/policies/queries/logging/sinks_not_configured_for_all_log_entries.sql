-- WITH found_sinks AS (
--     SELECT count(*) AS configured_sinks
--     FROM gcp_logging_sinks gls
--     WHERE gls.FILTER = ''
-- )
-- SELECT 'no sinks for all log entries configured' AS description
-- FROM found_sinks
-- WHERE configured_sinks = 0;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
WITH found_sinks AS (SELECT project_id, name, count(*) AS configured_sinks
                     FROM gcp_logging_sinks gls
                     WHERE gls.FILTER = ''
                     GROUP BY project_id, name)
SELECT "name"                                                             AS resource_id,
       :'execution_time'::timestamp                                       AS execution_time,
       :'framework'                                                       AS framework,
       :'check_id'                                                        AS check_id,
       'Ensure that sinks are configured for all log entries (Automated)' AS title,
       "project_id"                                                       AS project_id,
       CASE
           WHEN
               configured_sinks = 0
               THEN 'fail'
           ELSE 'pass'
           END                                                            AS status
FROM found_sinks;