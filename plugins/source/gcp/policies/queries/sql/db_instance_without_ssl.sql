-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
--     AND settings_ip_configuration_require_ssl = FALSE;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                               AS resource_id,
       :'execution_time'::timestamp                                                                           AS execution_time,
       :'framework'                                                                                           AS framework,
       :'check_id'                                                                                            AS check_id,
       'Ensure that the Cloud SQL database instance requires all incoming connections to use SSL (Automated)' AS title,
       gsi.project_id                                                                                         AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'SQLSERVER%'
                   AND (gsi.settings->'ipConfiguration'->>'requireSsl')::boolean = FALSE
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                AS status
FROM gcp_sql_instances gsi;