{% macro iam_avoid_root_usage(framework, check_id) %}
    select
        "name" as resource_id,
        "_cq_sync_time" as sync_time,
        '{{ framework }}' as framework,
        '{{ check_id }}' as check_id,
        'Ensure that the default network does not exist in a project (Automated)'
        as title,
        project_id as project_id,
        case when "name" = 'default' then 'fail' else 'pass' end as status
    from gcp_compute_networks
{% endmacro %}
