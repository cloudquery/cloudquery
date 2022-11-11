CREATE OR REPLACE VIEW pod_containers AS 
    SELECT
        uid,
	    container,
        pod.name as Name,
        pod.namespace as Namespace,
	    pod.context as Context
FROM k8s_core_pods pod
CROSS JOIN jsonb_array_elements(spec_containers) AS container;
