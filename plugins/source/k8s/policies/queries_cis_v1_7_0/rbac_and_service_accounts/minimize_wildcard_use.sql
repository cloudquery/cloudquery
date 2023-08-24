-- 5.1.3 Minimize wildcard use in Roles and ClusterRoles
INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Minimize wildcard use in Roles and ClusterRoles' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        	case
    when rule ->> 'apiGroups' like '%*%'
    or rule ->> 'resources' like '%*%'
    or rule ->> 'verbs' like '%*%' then 'fail'
    else 'pass'
  end as status
from
  k8s_rbac_cluster_roles,
  jsonb_array_elements(rules) rule
where
  name not like '%system%'
UNION
select
	uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Minimize wildcard use in Roles and ClusterRoles' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        case when
        rule ->> 'apiGroups' like '%*%'
        or rule ->> 'resources' like '%*%'
        or rule ->> 'verbs' like '%*%' then 'fail'
        else 'pass'
        end as status
from
  k8s_rbac_roles,
  jsonb_array_elements(rules) rule
where
  name not like '%system%'