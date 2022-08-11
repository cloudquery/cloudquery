INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select uid                                  AS resource_id,
       :'execution_time'::timestamp         AS execution_time,
       :'framework'                         AS framework,
       :'check_id'                          AS check_id,
       'Job containers privileges disabled' AS title,
       context                              AS context,
       namespace                            AS namespace,
       name                                 AS resource_name,
       CASE
           WHEN
               c -> 'securityContext' ->> 'privileged' = 'true'
               THEN 'fail'
           ELSE 'pass'
           END                              AS status
FROM k8s_batch_jobs,
     JSONB_ARRAY_ELEMENTS(template -> 'spec' -> 'containers') AS c