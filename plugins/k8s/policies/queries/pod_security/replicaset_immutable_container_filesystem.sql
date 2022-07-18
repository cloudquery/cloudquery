INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select uid                                                   AS resource_id,
       :'execution_time'::timestamp                          AS execution_time,
       :'framework'                                          AS framework,
       :'check_id'                                           AS check_id,
       'ReplicaSet containers root file system is read-only' AS title,
       context                                               AS context,
       namespace                                             AS namespace,
       name                                                  AS resource_name,
       CASE
           WHEN
               c -> 'securityContext' ->> 'readOnlyRootFilesystem' IS DISTINCT FROM 'true'
               THEN 'fail'
           ELSE 'pass'
           END                                               AS status
FROM k8s_apps_replica_sets,
     JSONB_ARRAY_ELEMENTS(template -> 'spec' -> 'containers') AS c;