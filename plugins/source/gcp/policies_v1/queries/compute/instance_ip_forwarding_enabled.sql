-- SELECT project_id, "name", self_link AS link
-- FROM gcp_compute_instances
-- WHERE can_ip_forward = TRUE;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                              AS resource_id,
       :'execution_time'::timestamp                                        AS execution_time,
       :'framework'                                                        AS framework,
       :'check_id'                                                         AS check_id,
       'Ensure that IP forwarding is not enabled on Instances (Automated)' AS title,
       project_id                                                          AS project_id,
       CASE
           WHEN
               can_ip_forward = TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                                             AS status
FROM gcp_compute_instances;