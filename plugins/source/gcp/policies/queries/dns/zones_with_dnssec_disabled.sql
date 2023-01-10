-- select * from gcp_dns_managed_zones
-- where visibility != 'private'
-- and ((dnssec_config is null) or (dnssec_config->>'state' = 'off'));


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "id"                                                      AS resource_id,
       :'execution_time'::timestamp                              AS execution_time,
       :'framework'                                              AS framework,
       :'check_id'                                               AS check_id,
       'Ensure that DNSSEC is enabled for Cloud DNS (Automated)' AS title,
       project_id                                                AS project_id,
       CASE
           WHEN
               visibility != 'private'
               and ((dnssec_config is null) or (dnssec_config->>'state' = 'off'))
               THEN 'fail'
           ELSE 'pass'
           END                                                   AS status
FROM gcp_dns_managed_zones;
