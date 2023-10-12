with
    project_policy_roles as (
        select _cq_sync_time, project_id, jsonb_array_elements(bindings) as binding
        from gcp_resourcemanager_project_policies
    ),
    role_members as (
        select
            _cq_sync_time,
            project_id,
            binding ->> 'role' as "role",
            jsonb_array_elements_text(binding -> 'members') as member
        from project_policy_roles
    )
select *
from role_members
