{% macro iam_service_account_admin_priv(framework, check_id) %}
SELECT member                                                            AS resource_id,
       _cq_sync_time                                                     AS sync_time,
       '{{framework}}'                                                      AS framework,
       '{{check_id}}'                                                       AS check_id,
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
FROM {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}