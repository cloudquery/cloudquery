CREATE OR REPLACE VIEW daemon_set_containers AS 
SELECT 
    uid,
	container,
    daemonset.name as Name,
    daemonset.namespace as Namespace,
	daemonset.context as Context
FROM k8s_apps_daemon_sets daemonset
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS container;