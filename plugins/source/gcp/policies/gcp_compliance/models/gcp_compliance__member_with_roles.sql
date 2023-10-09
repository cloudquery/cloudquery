SELECT
    _cq_sync_time,
    project_id,
    member,
    array_agg(role) AS roles
FROM {{ ref('gcp_compliance__project_policy_members') }}
GROUP BY _cq_sync_time, member, project_id