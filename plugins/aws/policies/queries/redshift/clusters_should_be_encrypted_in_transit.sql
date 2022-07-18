insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Connections to Amazon Redshift clusters should be encrypted in transit' as title,
    account_id,
    arn as resource_id,
    'fail' as status -- TODO FIXME
from aws_redshift_clusters as rsc

where exists(select 1
                    from aws_redshift_cluster_parameter_groups as rscpg
    inner join aws_redshift_cluster_parameters as rscp
        on
            rscpg.cq_id = rscp.cluster_parameter_group_cq_id
    where rsc.cq_id = rscpg.cluster_cq_id
        and (
            rscp.parameter_name = 'require_ssl' and rscp.parameter_value = 'false'
        )
        or (
            rscp.parameter_name = 'require_ssl' and rscp.parameter_value is null
        )
        or not exists((select 1
            from aws_redshift_cluster_parameters
            where cluster_parameter_group_cq_id = rscpg.cq_id
                and parameter_name = 'require_ssl'))
)
