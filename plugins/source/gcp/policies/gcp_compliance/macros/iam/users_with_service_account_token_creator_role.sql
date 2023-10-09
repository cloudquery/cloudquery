{% macro iam_users_with_service_account_token_creator_role(framework, check_id) %}
SELECT member                                                                                                                                AS resource_id,
       _cq_sync_time                                                                                                          AS sync_time,
       '{{framework}}'                                                                                                                          AS framework,
       '{{check_id}}'                                                                                                                           AS check_id,
       'Ensure that IAM users are not assigned the Service Account User or Service Account Token Creator roles at project level (Automated)' AS title,
       project_id                                                                                                                            AS project_id,
       CASE
           WHEN
                       "role" IN ('roles/iam.serviceAccountUser', 'roles/iam.serviceAccountTokenCreator')
                   AND "member" LIKE 'user:%'
               THEN 'fail'
           ELSE 'pass'
        END                                                                                                                               AS status
FROM {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}