-- SELECT project_id, name, self_link AS link
-- FROM gcp_storage_buckets
-- WHERE iam_configuration_uniform_bucket_level_access_enabled = FALSE;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                                   AS resource_id,
       :'execution_time'::timestamp                                                             AS execution_time,
       :'framework'                                                                             AS framework,
       :'check_id'                                                                              AS check_id,
       'Ensure that Cloud Storage buckets have uniform bucket-level access enabled (Automated)' AS title,
       project_id                                                                               AS project_id,
       CASE
           WHEN
               (uniform_bucket_level_access->>'Enabled')::boolean = FALSE
               THEN 'fail'
           ELSE 'pass'
           END                                                                                  AS status
FROM gcp_storage_buckets;