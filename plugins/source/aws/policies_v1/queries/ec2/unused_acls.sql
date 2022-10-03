with results as (
select distinct
    account_id,
    network_acl_id as resource_id,
    case when
        a->>'NetworkAclAssociationId' is null
        then 'pass'
        else 'fail'
    end as status
from aws_ec2_network_acls left join jsonb_array_elements(aws_ec2_network_acls.associations) as a on true
        )
insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Unused network access control lists should be removed' as title,
    account_id,
    resource_id,
    status
from results
