-- SELECT gsi.project_id, gsi.name, gsiia."type", gsi.self_link AS link
-- FROM gcp_sql_instances gsi
--     JOIN gcp_sql_instance_ip_addresses gsiia ON
--         gsi.cq_id = gsiia.instance_cq_id
-- WHERE database_version LIKE 'SQLSERVER%'
--     AND gsiia.type = 'PRIMARY' OR backend_type != 'SECOND_GEN';


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gsi.name                                                                      AS resource_id,
                :'execution_time'::timestamp                                                  AS execution_time,
                :'framework'                                                                  AS framework,
                :'check_id'                                                                   AS check_id,
                'Ensure that Cloud SQL database instances do not have public IPs (Automated)' AS title,
                gsi.project_id                                                                AS project_id,
                CASE
                    WHEN
                                    gsi.database_version LIKE 'SQLSERVER%'
                                AND gsiia.type = 'PRIMARY' OR gsi.backend_type != 'SECOND_GEN'
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                       AS status
FROM gcp_sql_instances gsi
         JOIN gcp_sql_instance_ip_addresses gsiia ON
    gsi.cq_id = gsiia.instance_cq_id;