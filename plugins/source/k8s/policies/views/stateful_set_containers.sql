CREATE OR REPLACE VIEW stateful_set_containers AS 
    SELECT
        uid,
	    container,
        stateful_set.name as Name,
        stateful_set.namespace as Namespace,
	    stateful_set.context as Context
FROM k8s_apps_stateful_sets stateful_set
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS container;
