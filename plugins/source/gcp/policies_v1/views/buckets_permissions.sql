CREATE OR REPLACE VIEW gcp_public_buckets_accesses AS
WITH project_policy_roles AS (SELECT p.project_id,
                                     p.name,
                                     CASE WHEN p.name like '%.appspot.com' THEN 'https://'||p.name ELSE 'https://' || p.name || '.storage.googleapis.com' END AS self_link,
                                     jsonb_array_elements(pp.bindings) AS binding
                              FROM gcp_storage_buckets p LEFT JOIN gcp_storage_bucket_policies pp ON pp.project_id=p.project_id AND pp.bucket_name=p.name
     ),
     role_members AS (SELECT project_id,
                             name,
                             self_link,
                             binding ->> 'role'                              AS "role",
                             jsonb_array_elements_text(binding -> 'members') AS MEMBER
                      FROM project_policy_roles)
SELECT project_id, name, self_link, "role", MEMBER
FROM role_members;
