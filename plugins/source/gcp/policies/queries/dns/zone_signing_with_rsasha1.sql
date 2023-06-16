INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gdmz.id                                                                                 AS resource_id,
                :'execution_time'::timestamp                                                            AS execution_time,
                :'framework'                                                                            AS framework,
                :'check_id'                                                                             AS check_id,
                'Ensure that RSASHA1 is not used for the zone-signing key in Cloud DNS DNSSEC (Manual)' AS title,
                gdmz.project_id                                                                         AS project_id,
                CASE
                    WHEN
                                gdmzdcdks->>'keyType' = 'zoneSigning'
                            AND gdmzdcdks->>'algorithm' = 'rsasha1'
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                 AS status
FROM gcp_dns_managed_zones gdmz, JSONB_ARRAY_ELEMENTS(gdmz.dnssec_config->'defaultKeySpecs') AS gdmzdcdks;