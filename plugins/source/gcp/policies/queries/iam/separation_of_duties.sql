INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT member                                                                                                            AS resource_id,
       :'execution_time'::timestamp                                                                                      AS execution_time,
       :'framework'                                                                                                      AS framework,
       :'check_id'                                                                                                       AS check_id,
       'Ensure that Separation of duties is enforced while assigning service account related roles to users (Automated)' AS title,
       project_id                                                                                                        AS project_id,
       CASE
           WHEN
                       "role" IN ('roles/iam.serviceAccountAdmin', 'roles/iam.serviceAccountUser')
                   AND "member" LIKE 'user:%'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                           AS status
FROM gcp_project_policy_members;
