INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gci.id                                                                     AS resource_id,
                :'execution_time'::timestamp                                               AS execution_time,
                :'framework'                                                               AS framework,
                :'check_id'                                                                AS check_id,
                'Ensure that Compute instances do not have public IP addresses (Automated' AS title,
                gci.project_id                                                             AS project_id,
                CASE
                    WHEN
                                gci."name" NOT LIKE 'gke-%'
                            AND (ac4->>'nat_i_p' IS NOT NULL OR ac4->>'nat_i_p' != '' OR ac6->>'nat_i_p' IS NOT NULL OR ac6->>'nat_i_p' != '')
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                    AS status
FROM gcp_compute_instances gci, JSONB_ARRAY_ELEMENTS(gci.network_interfaces) AS ni
LEFT JOIN JSONB_ARRAY_ELEMENTS(ni->'access_configs') AS ac4 ON TRUE
LEFT JOIN JSONB_ARRAY_ELEMENTS(ni->'ipv6_access_configs') AS ac6 ON TRUE
