-- SELECT
--     gls.project_id,
--     gls.name AS "sink_name",
--     gsb.name AS "bucket_name",
--     gsb.retention_policy_is_locked,
--     gsb.retention_policy_retention_period,
--     gls.destination
-- FROM gcp_logging_sinks gls
--     JOIN gcp_storage_buckets gsb ON
--         gsb.name = REPLACE(gls.destination, 'storage.googleapis.com/', '')
-- WHERE gls.destination LIKE 'storage.googleapis.com/%'
--     AND ( gsb.retention_policy_is_locked = FALSE
--         OR gsb.retention_policy_retention_period = 0)


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gsb.name                                                                                     AS resource_id,
                :'execution_time'::timestamp                                                                 AS execution_time,
                :'framework'                                                                                 AS framework,
                :'check_id'                                                                                  AS check_id,
                'Ensure that retention policies on log buckets are configured using Bucket Lock (Automated)' AS title,
                gls.project_id                                                                               AS project_id,
                CASE
                    WHEN
                                gls.destination LIKE 'storage.googleapis.com/%'
                            AND ((gsb.retention_policy->>'IsLocked')::boolean = FALSE
                            OR (gsb.retention_policy->>'RetentionPeriod')::integer = 0)
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                      AS status
FROM gcp_logging_sinks gls
         JOIN gcp_storage_buckets gsb ON
    gsb.name = REPLACE(gls.destination, 'storage.googleapis.com/', '');