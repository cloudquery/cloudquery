INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                       AS resource_id,
       :'execution_time'::timestamp                                                 AS execution_time,
       :'framework'                                                                 AS framework,
       :'check_id'                                                                  AS check_id,
       'Ensure Compute instances are launched with Shielded VM enabled (Automated)' AS title,
       project_id                                                                   AS project_id,
       CASE
           WHEN
               (shielded_instance_config->>'enable_integrity_monitoring')::boolean = FALSE
                   OR (shielded_instance_config->>'enable_vtpm')::boolean = FALSE
                   OR (shielded_instance_config->>'enable_secure_boot')::boolean = FALSE
               THEN 'fail'
           ELSE 'pass'
           END                                                                      AS status
FROM gcp_compute_instances;
