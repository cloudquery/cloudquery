-- SELECT project_id, name, self_link AS link
-- FROM gcp_compute_instances
-- WHERE metadata_items ->> 'block-project-ssh-keys' IS NULL
--     OR NOT metadata_items ->> 'block-project-ssh-keys' = ANY ('{1,true,True,TRUE,y,yes}');


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gci.name                                                                         AS resource_id,
       :'execution_time'::timestamp                                                     AS execution_time,
       :'framework'                                                                     AS framework,
       :'check_id'                                                                      AS check_id,
       'Ensure "Block Project-wide SSH keys" is enabled for VM instances (Automated)' AS title,
       gci.project_id                                                                       AS project_id,
       CASE
           WHEN
                gcmi->>'key' IS NULL OR
                NOT gcmi->>'value' = ANY ('{1,true,True,TRUE,y,yes}')
               THEN 'fail'
           ELSE 'pass'
           END                                                                          AS status
FROM gcp_compute_instances gci
    LEFT JOIN JSONB_ARRAY_ELEMENTS(gci.metadata->'items') gcmi ON gcmi->>'key' = 'block-project-ssh-keys';
