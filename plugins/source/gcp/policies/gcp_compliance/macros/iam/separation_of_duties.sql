{% macro iam_separation_of_duties(framework, check_id) %}
    select
        member as resource_id,
        _cq_sync_time as sync_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure that Separation of duties is enforced while assigning service account related roles to users (Automated)'
        as title,
        project_id as project_id,
        case
            when
                "role"
                in ('roles/iam.serviceAccountAdmin', 'roles/iam.serviceAccountUser')
                and "member" like 'user:%'
            then 'fail'
            else 'pass'
        end as status
    from {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}
