create table if not exists k8s_policy_results
(
    execution_time timestamp with time zone,
    framework      varchar(255),
    check_id       varchar(255),
    title          text,
    context        text,
    namespace      text,
    resource_id    varchar(1024),
    resource_name  text,
    status         varchar(16)
)