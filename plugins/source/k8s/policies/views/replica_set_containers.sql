CREATE OR REPLACE VIEW replica_set_containers AS
    SELECT
        uid,
	    container,
        replica_set.name as Name,
        replica_set.namespace as Namespace,
	    replica_set.context as Context
FROM k8s_apps_replica_sets replica_set
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS container;
