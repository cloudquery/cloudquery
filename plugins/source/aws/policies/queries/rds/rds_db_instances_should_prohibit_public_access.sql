INSERT INTO aws_policy_results
WITH internet_gateways AS (
    SELECT
        igw.internet_gateway_id,
        attachment->>'VpcId' AS vpc_id
    FROM aws_ec2_internet_gateways igw, jsonb_array_elements(igw.attachments) attachment
),
route_tables AS (
    SELECT
        route_table_id,
        association->>'SubnetId' AS subnet_id,
        route->>'DestinationCidrBlock' AS destination_cidr_block,
        route->>'GatewayId' AS gateway_id
    FROM aws_ec2_route_tables rt, jsonb_array_elements(rt.associations) association, jsonb_array_elements(rt.routes) route
),
failed_rds_instances AS (
    SELECT
        rds.arn
    FROM aws_rds_instances rds, jsonb_array_elements(rds.db_subnet_group->'Subnets') subnet
        JOIN aws_ec2_subnets ec2_subnet ON ec2_subnet.subnet_id = subnet->>'SubnetIdentifier'
        LEFT JOIN route_tables ON ec2_subnet.subnet_id = route_tables.subnet_id
        LEFT JOIN internet_gateways igw ON igw.internet_gateway_id = route_tables.gateway_id
    WHERE publicly_accessible
      AND destination_cidr_block = '0.0.0.0/0'
      AND igw.internet_gateway_id IS NOT NULL
    GROUP BY rds.arn
)
SELECT
    :'execution_time'                                                                                    AS execution_time,
    :'framework'                                                                                         AS framework,
    :'check_id'                                                                                          AS check_id,
    'RDS DB instances should prohibit public access, determined by the PubliclyAccessible configuration' AS title,
    account_id,
    aws_rds_instances.arn                                                                                AS resource_id,
    CASE
        WHEN failed_rds_instances.arn IS NOT NULL
        THEN 'fail'
        ELSE 'pass'
    END                                                                                                  AS status
FROM aws_rds_instances
    LEFT JOIN failed_rds_instances ON aws_rds_instances.arn = failed_rds_instances.arn
