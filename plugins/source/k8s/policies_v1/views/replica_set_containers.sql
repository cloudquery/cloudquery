CREATE OR REPLACE VIEW replica_set_containers AS
SELECT 
    uid,
    value AS container
FROM k8s_apps_replica_sets
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)
