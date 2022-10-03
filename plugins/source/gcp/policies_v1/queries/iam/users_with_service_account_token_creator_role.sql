-- SELECT project_id, "role", "member"
-- FROM gcp_project_policy_members
-- WHERE "role" IN ( 'roles/iam.serviceAccountUser', 'roles/iam.serviceAccountTokenCreator')
--     AND "member" LIKE 'user:%';


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT member                                                                                                                                AS resource_id,
       :'execution_time'::timestamp                                                                                                          AS execution_time,
       :'framework'                                                                                                                          AS framework,
       :'check_id'                                                                                                                           AS check_id,
       'Ensure that IAM users are not assigned the Service Account User or Service Account Token Creator roles at project level (Automated)' AS title,
       project_id                                                                                                                            AS project_id,
       CASE
           WHEN
                       "role" IN ('roles/iam.serviceAccountUser', 'roles/iam.serviceAccountTokenCreator')
                   AND "member" LIKE 'user:%'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                                               AS status
FROM gcp_project_policy_members;
