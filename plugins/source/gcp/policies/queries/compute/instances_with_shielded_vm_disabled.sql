-- SELECT project_id, gci."name", gci.self_link AS link
-- FROM gcp_compute_instances gci
-- WHERE shielded_instance_config_enable_integrity_monitoring = FALSE
--     OR shielded_instance_config_enable_vtpm = FALSE;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "id"                                                                         AS resource_id,
       :'execution_time'::timestamp                                                 AS execution_time,
       :'framework'                                                                 AS framework,
       :'check_id'                                                                  AS check_id,
       'Ensure Compute instances are launched with Shielded VM enabled (Automated)' AS title,
       project_id                                                                   AS project_id,
       CASE
           WHEN
               (shielded_instance_config->>'enable_integrity_monitoring')::boolean = FALSE
                   OR (shielded_instance_config->>'enable_vtpm')::boolean = FALSE
               THEN 'fail'
           ELSE 'pass'
           END                                                                      AS status
FROM gcp_compute_instances;
