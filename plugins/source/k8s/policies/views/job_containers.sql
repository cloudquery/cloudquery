CREATE OR REPLACE VIEW job_containers AS
    SELECT
        uid,
	    container,
        job.name as Name,
        job.namespace as Namespace,
	    job.context as Context
FROM k8s_batch_jobs job
CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS container;
