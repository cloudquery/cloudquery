-- SELECT gdmz.project_id, gdmz.id, gdmz.name, gdmz.dns_name
-- FROM gcp_dns_managed_zones gdmz
-- WHERE gdmz.dnssec_config_state != 'on'


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "id"                                                            AS resource_id,
       :'execution_time'::timestamp                                    AS execution_time,
       :'framework'                                                    AS framework,
       :'check_id'                                                     AS check_id,
       'Ensure legacy networks do not exist for a project (Automated)' AS title,
       project_id                                                      AS project_id,
       CASE
           WHEN
               dnssec_config->>'state' != 'on'
               THEN 'fail'
           ELSE 'pass'
           END                                                         AS status
FROM gcp_dns_managed_zones;