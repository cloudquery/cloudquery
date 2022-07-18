create table if not exists gcp_policy_results (
    execution_time time,
    framework varchar(255),
    check_id varchar(255),
    title text,
    project_id varchar(1024),
    resource_id varchar(1024),
    status varchar(16)
)