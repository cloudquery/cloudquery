-- SELECT
--     d.project_id,
--     d.id,
--     d.friendly_name,
--     d.self_link AS dataset_link,
--     a.special_group AS "group",
--     a."role"
-- FROM gcp_bigquery_datasets d
--     JOIN gcp_bigquery_dataset_accesses a ON
--         d.id = a.dataset_id
-- WHERE a."role" = 'allUsers'
--     OR a."role" = 'allAuthenticatedUsers';


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT d.id                                                                                   AS resource_id,
                :'execution_time'::timestamp                                                           AS execution_time,
                :'framework'                                                                           AS framework,
                :'check_id'                                                                            AS check_id,
                'Ensure that BigQuery datasets are not anonymously or publicly accessible (Automated)' AS title,
                d.project_id                                                                           AS project_id,
                CASE
                    WHEN
                                a->>'role' = 'allUsers'
                            OR a->>'role' = 'allAuthenticatedUsers'
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                AS status
FROM gcp_bigquery_datasets d, JSONB_ARRAY_ELEMENTS(d.access) AS a;
