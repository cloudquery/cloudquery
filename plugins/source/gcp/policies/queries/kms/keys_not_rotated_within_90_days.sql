-- SELECT *
-- FROM gcp_kms_keyring_crypto_keys gkkck
-- WHERE (rotation_period LIKE '%s'
--     AND REPLACE(rotation_period, 's', '')::NUMERIC > 7776000)
--     OR (rotation_period LIKE '%h'
--         AND REPLACE(rotation_period, 'h', '')::NUMERIC > 2160)
--     OR (rotation_period LIKE '%m'
--         AND REPLACE(rotation_period, 'm', '')::NUMERIC > 129600)
--     OR (rotation_period LIKE '%d'
--         AND REPLACE(rotation_period, 'd', '')::NUMERIC > 90)
--     OR DATE_PART('day', CURRENT_DATE - next_rotation_time ) > 90 ;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                          AS resource_id,
       :'execution_time'::timestamp                                                    AS execution_time,
       :'framework'                                                                    AS framework,
       :'check_id'                                                                     AS check_id,
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
FROM gcp_kms_crypto_keys;
