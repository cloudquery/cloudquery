CREATE OR REPLACE VIEW job_containers AS
    SELECT 
        uid,
        context,
        name,
        namespace,
        value AS container 
FROM k8s_batch_jobs
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value;
