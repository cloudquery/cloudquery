-- In the original document in CIS GCP v1.2.0, it describes the configuration should be 'off', but it is a typo.
-- This constraint has been updated on CIS GCP v1.3.0, this flag should be 'on'.
INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                                 AS resource_id,
       :'execution_time'::timestamp                                                                             AS execution_time,
       :'framework'                                                                                             AS framework,
       :'check_id'                                                                                              AS check_id,
       'Ensure "3625 (trace flag)" database flag for Cloud SQL SQL Server instance is set to "on" (Automated)'  AS title,
       gsi.project_id                                                                                           AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'SQLSERVER%'
                   AND (f->>'value' IS NULL
                   OR f->>'value' != 'on')
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                  AS status
FROM gcp_sql_instances gsi LEFT JOIN JSONB_ARRAY_ELEMENTS(gsi.settings->'databaseFlags') AS f ON f->>'name'='3625';
