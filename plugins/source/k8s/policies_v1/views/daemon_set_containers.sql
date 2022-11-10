CREATE OR REPLACE VIEW daemonset_containers AS 
SELECT 
    uid, value AS container
FROM k8s_apps_daemon_sets
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value