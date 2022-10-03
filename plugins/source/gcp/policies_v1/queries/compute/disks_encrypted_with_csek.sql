-- SELECT project_id, id, name, self_link AS link
-- FROM gcp_compute_disks
-- WHERE disk_encryption_key_sha256 IS NULL
--     OR disk_encryption_key_sha256 = ''
--     OR source_image_encryption_key_kms_key_name IS NULL
--     OR source_image_encryption_key_kms_key_name = '';


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "id"                                                                                                       AS resource_id,
       :'execution_time'::timestamp                                                                               AS execution_time,
       :'framework'                                                                                               AS framework,
       :'check_id'                                                                                                AS check_id,
       'Ensure VM disks for critical VMs are encrypted with Customer-Supplied Encryption Keys (CSEK) (Automated)' AS title,
       project_id                                                                                                 AS project_id,
       CASE
           WHEN
                   disk_encryption_key->>'sha256' IS NULL
                   OR disk_encryption_key->>'sha256' = ''
                   OR source_image_encryption_key->>'kms_key_name' IS NULL
                   OR source_image_encryption_key->>'kms_key_name' = ''
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                    AS status
FROM gcp_compute_disks;
