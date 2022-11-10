CREATE OR REPLACE VIEW pod_containers AS 
SELECT
    uid,
    value AS container 
FROM k8s_core_pods
CROSS JOIN jsonb_array_elements(spec_containers) AS value;
