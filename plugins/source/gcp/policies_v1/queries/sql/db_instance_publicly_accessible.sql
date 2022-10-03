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
                            AND gsisican->>'value' = '0.0.0.0/0'
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                          AS status
FROM gcp_sql_instances gsi, JSONB_ARRAY_ELEMENTS(gsi.settings->'ipConfiguration'->'authorizedNetworks') AS gsisican
