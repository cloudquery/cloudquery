{% macro iam_avoid_root_usage(framework, check_id) %}
  SELECT "name"                                                                   AS resource_id,
        "_cq_sync_time"                                                           AS sync_time,
        '{{ framework }}'                                                           AS framework,
        '{{ check_id }}'                                                            AS check_id,
        'Ensure that the default network does not exist in a project (Automated)' AS title,
        project_id                                                                AS project_id,
        CASE
            WHEN
                "name" = 'default'
                THEN 'fail'
            ELSE 'pass'
            END                                                                   AS status
  FROM gcp_compute_networks
{% endmacro %}