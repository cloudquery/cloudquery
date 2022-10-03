-- SELECT d.project_id, d.id, d.friendly_name, d.self_link AS dataset_link, t.self_link AS table_link
-- FROM gcp_bigquery_datasets d
--     JOIN gcp_bigquery_dataset_tables t ON
--         d.id = t.dataset_id
-- WHERE encryption_configuration_kms_key_name = '' OR default_encryption_configuration_kms_key_name IS NULL;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT d.id                                                                                                               AS resource_id,
                :'execution_time'::timestamp                                                                                       AS execution_time,
                :'framework'                                                                                                       AS framework,
                :'check_id'                                                                                                        AS check_id,
                'Ensure that a Default Customer-managed encryption key (CMEK) is specified for all BigQuery Data Sets (Automated)' AS title,
                d.project_id                                                                                                       AS project_id,
                CASE
                    WHEN
                                t.encryption_configuration->>'kmsKeyName' = '' OR
                                d.default_encryption_configuration->>'kmsKeyName' IS NULL -- TODO check if valid
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                                            AS status
FROM gcp_bigquery_datasets d
         JOIN gcp_bigquery_tables t ON
    d.dataset_reference->>'datasetId' = t.table_reference->>'datasetId' AND d.dataset_reference->>'projectId' = t.table_reference->>'projectId';
