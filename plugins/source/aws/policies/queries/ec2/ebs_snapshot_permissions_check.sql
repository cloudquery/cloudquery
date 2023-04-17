insert into aws_policy_results
WITH snapshot_access_groups AS (
    SELECT account_id,
           region,
           snapshot_id,
           JSONB_ARRAY_ELEMENTS(attribute->'CreateVolumePermissions') ->> 'Group' AS "group",
           JSONB_ARRAY_ELEMENTS(attribute->'CreateVolumePermissions') ->> 'UserId' AS user_id
    FROM aws_ec2_ebs_snapshots
)
SELECT DISTINCT
  :'execution_time'::timestamp as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Amazon EBS snapshots should not be public, determined by the ability to be restorable by anyone' as title,
  account_id,
  snapshot_id as resource_id,
  case when
    "group" = 'all'
    -- this is under question because
    -- trusted accounts(user_id) do not violate this control
        OR user_id IS DISTINCT FROM ''
    then 'fail'
    else 'pass'
  end as status
FROM snapshot_access_groups
