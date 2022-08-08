-- SELECT project_id, id, "name", self_link AS link
-- FROM gcp_compute_networks gcn
-- WHERE ip_v4_range IS NULL


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "id"                                                      AS resource_id,
       :'execution_time'::timestamp                              AS execution_time,
       :'framework'                                              AS framework,
       :'check_id'                                               AS check_id,
       'Ensure that DNSSEC is enabled for Cloud DNS (Automated)' AS title,
       project_id                                                AS project_id,
       CASE
           WHEN
               ip_v4_range IS NULL
               THEN 'fail'
           ELSE 'pass'
           END                                                   AS status
FROM gcp_compute_networks;