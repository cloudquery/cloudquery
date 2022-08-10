-- SELECT project_id, id, "name", self_link AS link
-- FROM gcp_compute_networks gcn
-- WHERE name = 'default';


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                    AS resource_id,
       :'execution_time'::timestamp                                              AS execution_time,
       :'framework'                                                              AS framework,
       :'check_id'                                                               AS check_id,
       'Ensure that the default network does not exist in a project (Automated)' AS title,
       project_id                                                                AS project_id,
       CASE
           WHEN
               "name" = 'default'
               THEN 'fail'
           ELSE 'pass'
           END                                                                   AS status
FROM gcp_compute_networks;