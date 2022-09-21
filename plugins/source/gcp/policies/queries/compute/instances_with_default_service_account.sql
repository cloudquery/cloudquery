-- SELECT project_id, gci."name", gci.self_link AS link
-- FROM gcp_compute_instances gci
--     JOIN gcp_compute_instance_service_accounts gcisa ON
--         gci.id = gcisa.instance_id
-- WHERE gci."name" NOT LIKE 'gke-'
--     AND gcisa.email = (SELECT default_service_account
--         FROM gcp_compute_projects
--         WHERE project_id = gci.project_id);


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gci.name                                                                                  AS resource_id,
                :'execution_time'::timestamp                                                              AS execution_time,
                :'framework'                                                                              AS framework,
                :'check_id'                                                                               AS check_id,
                'Ensure that instances are not configured to use the default service account (Automated)' AS title,
                gci.project_id                                                                            AS project_id,
                CASE
                    WHEN
                                gci."name" NOT LIKE 'gke-'
                            AND gcisa->>'email' = (SELECT default_service_account
                                               FROM gcp_compute_projects
                                               WHERE project_id = gci.project_id)
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                   AS status
FROM gcp_compute_instances gci, JSONB_ARRAY_ELEMENTS(gci.service_accounts) gcisa;