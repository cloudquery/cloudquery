-- SELECT project_id, id, friendly_name, self_link AS link
-- FROM gcp_bigquery_datasets
-- WHERE default_encryption_configuration_kms_key_name = ''
--     OR default_encryption_configuration_kms_key_name IS NULL;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT d.id                                                                                                      AS resource_id,
       :'execution_time'::timestamp                                                                            AS execution_time,
       :'framework'                                                                                            AS framework,
       :'check_id'                                                                                             AS check_id,
       'Ensure that all BigQuery Tables are encrypted with Customer-managed encryption key (CMEK) (Automated)' AS title,
       d.project_id                                                                                            AS project_id,
       CASE
           WHEN
                   d.default_encryption_configuration->>'kmsKeyName' = ''
                   OR d.default_encryption_configuration->>'kmsKeyName' IS NULL -- TODO check if valid
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                 AS status
FROM gcp_bigquery_datasets d;