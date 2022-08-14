INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select uid                                       AS resource_id,
       :'execution_time'::timestamp              AS execution_time,
       :'framework'                              AS framework,
       :'check_id'                               AS check_id,
       'DeamonSet containers to run as non-root' AS title,
       context                                   AS context,
       namespace                                 AS namespace,
       name                                      AS resource_name,
       CASE
           WHEN
               c -> 'securityContext' ->> 'runAsNonRoot' IS DISTINCT FROM 'true'
               THEN 'fail'
           ELSE 'pass'
           END                                   AS status
FROM k8s_apps_daemon_sets,
     JSONB_ARRAY_ELEMENTS(template -> 'spec' -> 'containers') AS c;