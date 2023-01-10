-- SELECT project_id, "role", "member"
-- FROM gcp_project_policy_members
-- WHERE ("role" IN ( 'roles/editor', 'roles/owner')
--     OR "role" LIKE ANY(ARRAY['%Admin', '%admin']))
--     AND "member" LIKE 'serviceAccount:%.iam.gserviceaccount.com';


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT member                                                            AS resource_id,
       :'execution_time'::timestamp                                      AS execution_time,
       :'framework'                                                      AS framework,
       :'check_id'                                                       AS check_id,
       'Ensure that Service Account has no Admin privileges (Automated)' AS title,
       project_id                                                        AS project_id,
       CASE
           WHEN
                   ("role" IN ('roles/editor', 'roles/owner') 
                   OR "role" LIKE ANY (ARRAY ['%Admin', '%admin']))
                   AND "member" LIKE 'serviceAccount:%.iam.gserviceaccount.com'
               THEN 'fail'
           ELSE 'pass'
           END                                                           AS status
FROM gcp_project_policy_members;
