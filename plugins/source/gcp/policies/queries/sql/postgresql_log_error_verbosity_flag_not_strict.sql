INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                                                AS resource_id,
       :'execution_time'::timestamp                                                                                            AS execution_time,
       :'framework'                                                                                                            AS framework,
       :'check_id'                                                                                                             AS check_id,
       'Ensure "log_error_verbosity" database flag for Cloud SQL PostgreSQL instance is set to "DEFAULT" or stricter (Manual)' AS title,
       gsi.project_id                                                                                                          AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'POSTGRES%'
                   AND (f->>'value' IS NULL
                   OR f->>'value' NOT IN ('default', 'terse'))
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                                 AS status
FROM gcp_sql_instances gsi LEFT JOIN JSONB_ARRAY_ELEMENTS(gsi.settings->'databaseFlags') AS f ON f->>'name'='log_error_verbosity';
