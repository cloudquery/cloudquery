insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure a log metric filter and alarm exist for VPC changes (Scored)' as title,
  account_id,
  cloud_watch_logs_log_group_arn as resource_id,
  case when pattern NOT LIKE '%NOT%'
           AND pattern LIKE '%($.eventName = CreateVpc)%'
           AND pattern LIKE '%($.eventName = DeleteVpc)%'
           AND pattern LIKE '%($.eventName = ModifyVpcAttribute)%'
           AND pattern LIKE '%($.eventName = AcceptVpcPeeringConnection)%'
           AND pattern LIKE '%($.eventName = CreateVpcPeeringConnection)%'
           AND pattern LIKE '%($.eventName = DeleteVpcPeeringConnection)%'
           AND pattern LIKE '%($.eventName = RejectVpcPeeringConnection)%'
           AND pattern LIKE '%($.eventName = AttachClassicLinkVpc)%'
           AND pattern LIKE '%($.eventName = DetachClassicLinkVpc)%'
           AND pattern LIKE '%($.eventName = DisableVpcClassicLink)%'
           AND pattern LIKE '%($.eventName = EnableVpcClassicLink)%'
      then 'pass'
      else 'fail'
  end as status
from view_aws_log_metric_filter_and_alarm
