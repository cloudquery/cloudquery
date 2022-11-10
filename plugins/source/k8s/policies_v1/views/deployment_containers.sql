CREATE OR REPLACE VIEW deployment_containers AS 
    SELECT
        uid,
        value AS container 
FROM k8s_apps_deployments
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value;