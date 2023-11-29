create table if not exists aws_policy_results (
    execution_time timestamp with time zone,
    framework varchar(255),
    check_id varchar(255),
    title text,
    account_id varchar(1024),
    resource_id varchar(1024),
    status varchar(16)
)