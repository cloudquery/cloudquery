CREATE OR REPLACE VIEW gcp_project_policy_members AS
WITH project_policy_roles AS (
    SELECT project_id,
        jsonb_array_elements(p.policy -> 'bindings') AS binding
    FROM gcp_resource_manager_projects p
),
role_members AS (
    SELECT
        project_id,
        binding ->> 'role' AS "role",
        jsonb_array_elements_text(binding -> 'members') AS MEMBER
    FROM project_policy_roles
)
SELECT *
FROM role_members;
