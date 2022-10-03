-- SELECT *
-- FROM gcp_firewall_allowed_rules
-- WHERE direction = 'INGRESS'
--     AND ( ip_protocol = 'tcp'
--         OR ip_protocol = 'all' )
--     AND '0.0.0.0/0' = ANY(source_ranges)
--     AND (3986 BETWEEN range_start AND range_end
--         OR '3986' = single_port
--         OR CARDINALITY(ports) = 0
--         OR ports IS NULL)


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                               AS resource_id,
       :'execution_time'::timestamp                                         AS execution_time,
       :'framework'                                                         AS framework,
       :'check_id'                                                          AS check_id,
       'Ensure that RDP access is restricted from the Internet (Automated)' AS title,
       project_id                                                           AS project_id,
       CASE
           WHEN
                       direction = 'INGRESS'
                   AND (ip_protocol = 'tcp'
                   OR ip_protocol = 'all')
                   AND '0.0.0.0/0' = ANY (source_ranges)
                   AND (3986 BETWEEN range_start AND range_end
                   OR '3986' = single_port
                   OR CARDINALITY(ports) = 0
                   OR ports IS NULL)
               THEN 'fail'
           ELSE 'pass'
           END                                                              AS status
FROM gcp_firewall_allowed_rules;