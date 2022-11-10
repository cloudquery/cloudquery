CREATE OR REPLACE VIEW daemon_set_containers AS 
SELECT 
    uid, 
    daemonset.context as 'Context',
    daemonset.name as 'Name',
    daemonset.namespace as 'Namespace',
    value AS container
FROM k8s_apps_daemon_sets daemonset 
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value;