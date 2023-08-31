INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "name"                                                                                   AS resource_id,
       :'execution_time'::timestamp                                                             AS execution_time,
       :'framework'                                                                             AS framework,
       :'check_id'                                                                              AS check_id,
       'Ensure that Cloud Storage bucket is not anonymously or publicly accessible (Automated)' AS title,
       project_id                                                                               AS project_id,
       CASE
           WHEN
                       member LIKE '%allUsers%'
                   OR member LIKE '%allAuthenticatedUsers%'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                  AS status
FROM gcp_public_buckets_accesses;
