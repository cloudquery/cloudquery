WITH project_policy_roles AS (
    SELECT 
        _cq_sync_time,
        project_id,
        jsonb_array_elements(bindings) AS binding
    FROM gcp_resourcemanager_project_policies
),
role_members AS (
    SELECT
        _cq_sync_time,
        project_id,
        binding ->> 'role' AS "role",
        jsonb_array_elements_text(binding -> 'members') AS MEMBER
    FROM project_policy_roles
)
SELECT *
FROM role_members