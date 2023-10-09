select _cq_sync_time, project_id, member, array_agg(role) as roles
from {{ ref('gcp_compliance__project_policy_members') }}
group by _cq_sync_time, member, project_id
