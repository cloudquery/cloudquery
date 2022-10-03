-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
--     AND (settings_database_flags IS NULL
--         OR settings_database_flags ->> 'log_checkpoints' != 'on'
--         OR settings_database_flags ->> 'log_checkpoints' IS NULL);


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                                       AS resource_id,
       :'execution_time'::timestamp                                                                                   AS execution_time,
       :'framework'                                                                                                   AS framework,
       :'check_id'                                                                                                    AS check_id,
       'Ensure that the "log_checkpoints" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)' AS title,
       gsi.project_id                                                                                                 AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'POSTGRES%'
                   AND (f->>'value' IS NULL
                   OR f->>'value' != 'on')
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                        AS status
FROM gcp_sql_instances gsi LEFT JOIN JSONB_ARRAY_ELEMENTS(gsi.settings->'databaseFlags') AS f ON f->>'name'='log_checkpoints';
