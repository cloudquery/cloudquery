{% macro kms_publicly_accessible(framework, check_id) %}
SELECT member                                                                                    AS resource_id,
       _cq_sync_time                                                              AS sync_time,
       '{{framework}}'                                                                              AS framework,
       '{{check_id}}'                                                                               AS check_id,
       'Ensure that Cloud KMS cryptokeys are not anonymously or publicly accessible (Automated)' AS title,
       project_id                                                                                AS project_id,
       CASE
           WHEN
                       "member" LIKE '%allUsers%'
                   OR "member" LIKE '%allAuthenticatedUsers%'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                   AS status
FROM {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}