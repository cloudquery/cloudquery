-- SELECT project_id, name, self_link AS link
-- FROM gcp_compute_instances
-- WHERE metadata_items ->> 'serial-port-enable' = ANY ('{1,true,True,TRUE,y,yes}');


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                           AS resource_id,
       :'execution_time'::timestamp                                                     AS execution_time,
       :'framework'                                                                     AS framework,
       :'check_id'                                                                      AS check_id,
       'Ensure "Enable connecting to serial ports" is not enabled for VM Instance (Automated)' AS title,
       project_id                                                                       AS project_id,
       CASE
           WHEN
             gcmi->>'key' = 'serial-port-enable' AND gcmi->>'value' = ANY ('{1,true,True,TRUE,y,yes}')
               THEN 'fail'
           ELSE 'pass'
           END                                                                          AS status
FROM gcp_compute_instances gci, JSONB_ARRAY_ELEMENTS(gci.metadata->'items') gcmi;
