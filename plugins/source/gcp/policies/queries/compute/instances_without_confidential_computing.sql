-- SELECT project_id, "name", gci.self_link AS link
-- FROM gcp_compute_instances gci
-- WHERE confidential_instance_config_enable_confidential_compute = FALSE;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                                  AS resource_id,
       :'execution_time'::timestamp                                                            AS execution_time,
       :'framework'                                                                            AS framework,
       :'check_id'                                                                             AS check_id,
       'Ensure that Compute instances have Confidential Computing enabled (Automated)' AS title,
       project_id                                                                              AS project_id,
       CASE
           WHEN
               (confidential_instance_config->>'enable_confidential_compute')::boolean = FALSE
               THEN 'fail'
           ELSE 'pass'
           END                                                                                 AS status
FROM gcp_compute_instances;