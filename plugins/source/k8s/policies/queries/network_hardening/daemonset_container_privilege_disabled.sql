INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
      :'execution_time'::timestamp      AS execution_time,
      :'framework'                      AS framework,
      :'check_id'                       AS check_id,
      'Daemonset container privilege disabled' AS title,
      context                           AS context,
      namespace                         AS namespace,
      name                              AS resource_name,
      CASE WHEN
          (SELECT COUNT(*) FROM daemonset_containers 
            WHERE daemon_set_containers.uid = k8s_apps_daemon_sets.uid 
            AND daemon_set_containers.container->'securityContext'->>'privileged' = 'true') > 0
          THEN 'fail'
          ELSE 'pass'
          END                          AS status
FROM k8s_apps_daemon_sets;
