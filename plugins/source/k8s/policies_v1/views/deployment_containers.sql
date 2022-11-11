CREATE OR REPLACE VIEW deployment_containers AS 
    SELECT
        uid,
	    container,
        deployment.name as Name,
        deployment.namespace as Namespace,
	    deployment.context as Context
FROM k8s_apps_deployments deployment
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS container;