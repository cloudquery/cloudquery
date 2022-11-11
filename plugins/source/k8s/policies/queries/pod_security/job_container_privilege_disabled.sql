WITH job_containers AS (SELECT uid, value AS container 
                        FROM k8s_batch_jobs
                        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                                         AS resource_id,
       :'execution_time'::timestamp         AS execution_time,
       :'framework'                         AS framework,
       :'check_id'                          AS check_id,
       'Job containers privileges disabled' AS title,
       context                              AS context,
       namespace                            AS namespace,
       name                                 AS resource_name,
       CASE
           WHEN
               (SELECT COUNT(*) FROM job_containers WHERE job_containers.uid = k8s_batch_jobs.uid AND
              job_containers.container->'securityContext'->>'privileged' = 'true') > 0
               THEN 'fail'
           ELSE 'pass'
           END                              AS status
FROM k8s_batch_jobs

