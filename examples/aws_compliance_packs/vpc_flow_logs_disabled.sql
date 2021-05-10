SELECT aws_ec2_vpcs.* FROM aws_ec2_vpcs
                               LEFT JOIN aws_ec2_flow_logs ON aws_ec2_flow_logs.resource_id = aws_ec2_vpcs.vpc_id
WHERE aws_ec2_flow_logs.resource_id IS NULL OR
        aws_ec2_flow_logs.flow_log_status != 'ACTIVE' OR
        aws_ec2_flow_logs.traffic_type != 'ALL';