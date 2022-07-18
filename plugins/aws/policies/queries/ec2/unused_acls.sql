insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Unused network access control lists should be removed' as title,
    account_id,
    id as resource_id,
    case when
        aws_ec2_network_acl_associations.cq_id is null
        then 'pass'
        else 'fail'
    end as status
from aws_ec2_network_acls
     left join
    aws_ec2_network_acl_associations on
        aws_ec2_network_acls.cq_id = aws_ec2_network_acl_associations.network_acl_cq_id
