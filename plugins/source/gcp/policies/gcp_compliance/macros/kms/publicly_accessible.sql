{% macro kms_publicly_accessible(framework, check_id) %}
    select
        member as resource_id,
        _cq_sync_time as sync_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure that Cloud KMS cryptokeys are not anonymously or publicly accessible (Automated)'
        as title,
        project_id as project_id,
        case
            when "member" like '%allUsers%' or "member" like '%allAuthenticatedUsers%'
            then 'fail'
            else 'pass'
        end as status
    from {{ ref('gcp_compliance__project_policy_members') }}
{% endmacro %}
