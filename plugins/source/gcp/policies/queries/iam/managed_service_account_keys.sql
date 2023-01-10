INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gisa.name                                                                                          AS resource_id,
                :'execution_time'::timestamp                                                                       AS execution_time,
                :'framework'                                                                                       AS framework,
                :'check_id'                                                                                        AS check_id,
                'Ensure that there are only GCP-managed service account keys for each service account (Automated)' AS title,
                gisa.project_id                                                                                    AS project_id,
                CASE
                    WHEN
                                gisa.email LIKE '%iam.gserviceaccount.com' AND gisak."key_type" = 'USER_MANAGED'
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                            AS status
FROM gcp_iam_service_accounts gisa
         JOIN gcp_iam_service_account_keys gisak ON
    gisa.project_id = gisak.project_id AND gisa.unique_id = gisak.service_account_unique_id;
