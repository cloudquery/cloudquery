{% macro kms_keys_not_rotated_within_90_days(framework, check_id) %}
SELECT "name"                                                                          AS resource_id,
       _cq_sync_time                                                    AS sync_time,
       '{{framework}}'                                                                    AS framework,
       '{{check_id}}'                                                                     AS check_id,
       'Ensure KMS encryption keys are rotated within a period of 90 days (Automated)' AS title,
       project_id                                                                      AS project_id,
       CASE
           WHEN
               (make_interval(secs => rotation_period/1000000000.0) > make_interval(days => 90))
               OR next_rotation_time IS NULL
               OR DATE_PART('day', CURRENT_DATE - next_rotation_time::timestamp) > 90
               THEN 'fail'
           ELSE 'pass'
           END                                                                         AS status
FROM gcp_kms_crypto_keys
{% endmacro %}