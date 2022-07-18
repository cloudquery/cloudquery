-- SELECT gsi.project_id, gsi.name, gsisican.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
--     JOIN gcp_sql_instance_settings_ip_config_authorized_networks gsisican ON
--         gsi.cq_id = gsisican.instance_cq_id
-- WHERE database_version LIKE 'SQLSERVER%'
--     AND gsisican.value = '0.0.0.0/0'


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gsi.name                                                                         AS resource_id,
                :'execution_time'::timestamp                                                     AS execution_time,
                :'framework'                                                                     AS framework,
                :'check_id'                                                                      AS check_id,
                'Ensure that Cloud SQL database instances are not open to the world (Automated)' AS title,
                gsi.project_id                                                                   AS project_id,
                CASE
                    WHEN
                                gsi.database_version LIKE 'SQLSERVER%'
                            AND gsisican.value = '0.0.0.0/0'
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                          AS status
FROM gcp_sql_instances gsi
         JOIN gcp_sql_instance_settings_ip_config_authorized_networks gsisican ON
    gsi.cq_id = gsisican.instance_cq_id;