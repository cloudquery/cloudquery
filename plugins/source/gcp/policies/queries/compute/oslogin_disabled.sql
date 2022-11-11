-- SELECT project_id, name, self_link AS link
-- FROM gcp_compute_projects
-- WHERE common_instance_metadata_items ->> 'enable-oslogin' IS NULL
--     OR NOT common_instance_metadata_items ->> 'enable-oslogin' = ANY ('{1,true,True,TRUE,y,yes}');


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                AS resource_id,
       :'execution_time'::timestamp                          AS execution_time,
       :'framework'                                          AS framework,
       :'check_id'                                           AS check_id,
       'Ensure oslogin is enabled for a Project (Automated)' AS title,
       project_id                                            AS project_id,
       CASE
           WHEN
               cimd->>'key' IS NULL OR
               NOT cimd->>'value' = ANY ('{1,true,True,TRUE,y,yes}')
               THEN 'fail'
           ELSE 'pass'
           END                                               AS status
FROM gcp_compute_projects
    LEFT JOIN JSONB_ARRAY_ELEMENTS(common_instance_metadata->'items') cimd ON cimd->>'key' = 'enable-oslogin';
