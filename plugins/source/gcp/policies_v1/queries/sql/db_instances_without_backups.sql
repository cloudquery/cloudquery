-- SELECT project_id, name, self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
--     AND settings_backup_enabled = FALSE;

INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                     AS resource_id,
       :'execution_time'::timestamp                                                                 AS execution_time,
       :'framework'                                                                                 AS framework,
       :'check_id'                                                                                  AS check_id,
       'Ensure that Cloud SQL database instances are configured with automated backups (Automated)' AS title,
       gsi.project_id                                                                               AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'SQLSERVER%'
                   AND (gsi.settings->'backupConfiguration'->>'enabled')::boolean = FALSE
               THEN 'fail'
           ELSE 'pass'
           END                                                                                      AS status
FROM gcp_sql_instances gsi;