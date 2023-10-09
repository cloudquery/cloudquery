{% macro iam_separation_of_duties(framework, check_id) %}
SELECT member                                                                                                            AS resource_id,
       _cq_sync_time                                                                                                     AS sync_time,
       '{{framework}}'                                                                                                      AS framework,
       '{{check_id}}'                                                                                                       AS check_id,
       'Ensure that Separation of duties is enforced while assigning service account related roles to users (Automated)' AS title,
       project_id                                                                                                        AS project_id,
       CASE
           WHEN
                       "role" IN ('roles/iam.serviceAccountAdmin', 'roles/iam.serviceAccountUser')
                   AND "member" LIKE 'user:%'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                           AS status
FROM {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}