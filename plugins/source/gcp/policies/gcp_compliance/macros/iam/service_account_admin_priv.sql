{% macro iam_service_account_admin_priv(framework, check_id) %}
    select
        member as resource_id,
        _cq_sync_time as sync_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure that Service Account has no Admin privileges (Automated)' as title,
        project_id as project_id,
        case
            when
                (
                    "role" in ('roles/editor', 'roles/owner')
                    or "role" like any (array['%Admin', '%admin'])
                )
                and "member" like 'serviceAccount:%.iam.gserviceaccount.com'
            then 'fail'
            else 'pass'
        end as status
    from {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}
