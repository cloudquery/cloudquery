CREATE OR REPLACE VIEW gcp_public_buckets_accesses AS
WITH project_policy_roles AS (SELECT project_id,
                                     name,
                                     self_link,
                                     jsonb_array_elements(p.policy -> 'bindings') AS binding
                              FROM gcp_storage_buckets p),
     role_members AS (SELECT project_id,
                             name,
                             self_link,
                             binding ->> 'role'                              AS "role",
                             jsonb_array_elements_text(binding -> 'members') AS MEMBER
                      FROM project_policy_roles)
SELECT project_id, name, self_link, "role", MEMBER
FROM role_members
