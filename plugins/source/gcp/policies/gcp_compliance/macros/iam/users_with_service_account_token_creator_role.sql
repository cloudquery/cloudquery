{% macro iam_users_with_service_account_token_creator_role(framework, check_id) %}
    select
        member as resource_id,
        _cq_sync_time as sync_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure that IAM users are not assigned the Service Account User or Service Account Token Creator roles at project level (Automated)'
        as title,
        project_id as project_id,
        case
            when
                "role" in (
                    'roles/iam.serviceAccountUser',
                    'roles/iam.serviceAccountTokenCreator'
                )
                and "member" like 'user:%'
            then 'fail'
            else 'pass'
        end as status
    from {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}
