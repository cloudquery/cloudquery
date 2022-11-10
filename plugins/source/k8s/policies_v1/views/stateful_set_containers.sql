CREATE OR REPLACE VIEW stateful_set_containers AS 
SELECT 
    uid,
    context,
    name,
    namespace,
    value AS container
FROM k8s_apps_stateful_sets
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value;
