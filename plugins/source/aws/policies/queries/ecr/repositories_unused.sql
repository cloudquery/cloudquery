insert into aws_policy_results
with image as (select distinct repository_cq_id from aws_ecr_repository_images)
select :'execution_time'       as execution_time,
       :'framework'            as framework,
       :'check_id'             as check_id,
       'Unused ECR repository' as title,
       repository.account_id,
       repository.arn          as resource_id,
       'fail'                  as status
from aws_ecr_repositories repository
         left join image on image.repository_cq_id = repository.cq_id
where image.repository_cq_id is null