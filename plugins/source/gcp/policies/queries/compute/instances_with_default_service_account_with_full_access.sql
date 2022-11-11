-- SELECT *
-- FROM gcp_compute_instances gci
--     JOIN gcp_compute_instance_service_accounts gcisa ON
--         gci.id = gcisa.instance_id
-- WHERE gcisa.email = (SELECT default_service_account
--     FROM gcp_compute_projects
--     WHERE project_id = gci.project_id)
--     AND 'https://www.googleapis.com/auth/cloud-platform' = ANY(gcisa.scopes);


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gci.name                                                                                                                     AS resource_id,
                :'execution_time'::timestamp                                                                                                 AS execution_time,
                :'framework'                                                                                                                 AS framework,
                :'check_id'                                                                                                                  AS check_id,
                'Ensure that instances are not configured to use the default service account with full access to all Cloud APIs (Automated)' AS title,
                gci.project_id                                                                                                               AS project_id,
                CASE
                    WHEN
                            gcisa->>'email' = (SELECT default_service_account
                                               FROM gcp_compute_projects
                                               WHERE project_id = gci.project_id)
                            AND ARRAY['https://www.googleapis.com/auth/cloud-platform'] <@ ARRAY(SELECT JSONB_ARRAY_ELEMENTS_TEXT(gcisa->'scopes'))
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                                                      AS status
FROM gcp_compute_instances gci, JSONB_ARRAY_ELEMENTS(gci.service_accounts) gcisa;