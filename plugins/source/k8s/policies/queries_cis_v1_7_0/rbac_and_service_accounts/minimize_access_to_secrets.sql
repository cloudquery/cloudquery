--5.1.2 Minimize access to secrets
INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Minimize access to secrets' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        case when
            rule -> 'resources' ? 'secrets' and rule -> 'verbs'  @> '["get", "list", "watch"]'::jsonb
            then 'fail'
            else 'pass'
        end as status
from
  k8s_rbac_cluster_roles,
  jsonb_array_elements(rules) rule
where
  name not like '%system%'