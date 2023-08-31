INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                                 AS resource_id,
       :'execution_time'::timestamp                                                                             AS execution_time,
       :'framework'                                                                                             AS framework,
       :'check_id'                                                                                              AS check_id,
       'Ensure "log_planner_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)' AS title,
       gsi.project_id                                                                                           AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'POSTGRES%'
                   AND (f->>'value' IS NULL
                   OR f->>'value' != 'off')
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                  AS status
FROM gcp_sql_instances gsi LEFT JOIN JSONB_ARRAY_ELEMENTS(gsi.settings->'databaseFlags') AS f ON f->>'name'='log_planner_stats';
