-- aws_apigateway_api_keys
ALTER TABLE aws_apigateway_api_keys ADD COLUMN arn TEXT;

UPDATE aws_apigateway_api_keys
SET arn = format('arn:aws:apigateway:%s::/apikeys/%s', region, id);


-- aws_apigateway_client_certificates
ALTER TABLE aws_apigateway_client_certificates ADD COLUMN arn TEXT;

UPDATE aws_apigateway_client_certificates
SET arn = format('arn:aws:apigateway:%s::/clientcertificates/%s', region, id);


-- aws_apigateway_domain_names
ALTER TABLE aws_apigateway_domain_names ADD COLUMN arn TEXT;

UPDATE aws_apigateway_domain_names
SET arn = format('arn:aws:apigateway:%s::/domainnames/%s', region, domain_name);

ALTER TABLE aws_apigateway_domain_name_base_path_mappings ADD COLUMN arn TEXT;

UPDATE aws_apigateway_domain_name_base_path_mappings t
SET arn = format('arn:aws:apigateway:%s::/domainnames/%s/basepathmappings/%s', d.region, d.domain_name, t.base_path)
FROM aws_apigateway_domain_names d
WHERE d.cq_id = t.domain_name_cq_id;


-- aws_apigateway_rest_apis
ALTER TABLE aws_apigateway_rest_apis ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_apis
SET arn = format('arn:aws:apigateway:%s::/restapis/%s', region, id);

ALTER TABLE aws_apigateway_rest_api_authorizers ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_authorizers t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/authorizers/%s', region, a.id, t.id)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_deployments ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_deployments t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/deployments/%s', region, a.id, t.id)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_documentation_parts ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_documentation_parts t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/documentation/parts/%s', region, a.id, t.id)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_documentation_versions ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_documentation_versions t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/documentation/versions/%s', region, a.id, t.version)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_gateway_responses ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_gateway_responses t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/gatewayresponses/%s', region, a.id, t.response_type)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_models ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_models t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/models/%s', region, a.id, t.name)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_request_validators ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_request_validators t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/requestvalidators/%s', region, a.id, t.id)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_resources ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_resources t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/resources/%s', region, a.id, t.id)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;

ALTER TABLE aws_apigateway_rest_api_stages ADD COLUMN arn TEXT;

UPDATE aws_apigateway_rest_api_stages t
SET arn = format('arn:aws:apigateway:%s::/restapis/%s/stages/%s', region, a.id, t.stage_name)
FROM aws_apigateway_rest_apis a
WHERE a.cq_id = t.rest_api_cq_id;


-- aws_apigateway_usage_plans
ALTER TABLE aws_apigateway_usage_plans ADD COLUMN arn TEXT;

UPDATE aws_apigateway_usage_plans
SET arn = format('arn:aws:apigateway:%s::/usageplans/%s', region, id);

ALTER TABLE aws_apigateway_usage_plan_keys ADD COLUMN arn TEXT;

UPDATE aws_apigateway_usage_plan_keys t
SET arn = format('arn:aws:apigateway:%s::/usageplans/%s/keys/%s', region, p.id, t.id)
FROM aws_apigateway_usage_plans p
WHERE p.cq_id = t.usage_plan_cq_id;


-- aws_apigateway_vpc_links
ALTER TABLE aws_apigateway_vpc_links ADD COLUMN arn TEXT;

UPDATE aws_apigateway_vpc_links
SET arn = format('arn:aws:apigateway:%s::/vpclinks/%s', region, id);


-- aws_apigatewayv2_apis
ALTER TABLE aws_apigatewayv2_apis ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_apis
SET arn = format('arn:aws:apigateway:%s::/apis/%s', region, id);

ALTER TABLE aws_apigatewayv2_api_authorizers ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_authorizers t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/authorizers/%s', region, a.id, t.authorizer_id)
FROM aws_apigatewayv2_apis a
WHERE a.cq_id = t.api_cq_id;

ALTER TABLE aws_apigatewayv2_api_deployments ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_deployments t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/deployments/%s', region, a.id, t.deployment_id)
FROM aws_apigatewayv2_apis a
WHERE a.cq_id = t.api_cq_id;

ALTER TABLE aws_apigatewayv2_api_integrations ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_integrations t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/integrations/%s', region, a.id, t.integration_id)
FROM aws_apigatewayv2_apis a
WHERE a.cq_id = t.api_cq_id;

ALTER TABLE aws_apigatewayv2_api_integration_responses ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_integration_responses t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/integrationresponses/%s', region, a.id, t.integration_response_id)
FROM aws_apigatewayv2_apis a, aws_apigatewayv2_api_integrations b
WHERE a.cq_id = b.api_cq_id AND b.cq_id = t.api_integration_cq_id;

ALTER TABLE aws_apigatewayv2_api_models ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_models t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/models/%s', region, a.id, t.model_id)
FROM aws_apigatewayv2_apis a
WHERE a.cq_id = t.api_cq_id;

ALTER TABLE aws_apigatewayv2_api_routes ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_routes t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/routes/%s', region, a.id, t.route_id)
FROM aws_apigatewayv2_apis a
WHERE a.cq_id = t.api_cq_id;

ALTER TABLE aws_apigatewayv2_api_route_responses ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_route_responses t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/routes/%s/routeresponses/%s', region, a.id, b.route_id, t.route_response_id)
FROM aws_apigatewayv2_apis a, aws_apigatewayv2_api_routes b
WHERE a.cq_id = b.api_cq_id AND b.cq_id = t.api_route_cq_id;

ALTER TABLE aws_apigatewayv2_api_stages ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_api_stages t
SET arn = format('arn:aws:apigateway:%s::/apis/%s/stages/%s', region, a.id, t.stage_name)
FROM aws_apigatewayv2_apis a
WHERE a.cq_id = t.api_cq_id;


-- aws_apigatewayv2_domain_names
ALTER TABLE aws_apigatewayv2_domain_names ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_domain_names
SET arn = format('arn:aws:apigateway:%s::/domainnames/%s', region, domain_name);

ALTER TABLE aws_apigatewayv2_domain_name_rest_api_mappings ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_domain_name_rest_api_mappings t
SET arn = format('arn:aws:apigateway:%s::/domainnames/%s/apimappings/%s', a.region, a.domain_name, t.api_mapping_id)
FROM aws_apigatewayv2_domain_names a
WHERE a.cq_id = t.domain_name_cq_id;


-- aws_apigatewayv2_vpc_links
ALTER TABLE aws_apigatewayv2_vpc_links ADD COLUMN arn TEXT;

UPDATE aws_apigatewayv2_vpc_links
SET arn = format('arn:aws:apigateway:%s::/vpclinks/%s', region, id);


-- aws_cloudfront_cache_policies
ALTER TABLE aws_cloudfront_cache_policies ADD COLUMN arn TEXT;

UPDATE aws_cloudfront_cache_policies
SET arn = format('arn:aws:cloudfront::%s:cache-policy/%s', account_id, id);


-- aws_cognito_identity_pools
ALTER TABLE aws_cognito_identity_pools ADD COLUMN arn TEXT;

UPDATE aws_cognito_identity_pools
SET arn = format('arn:aws:cognito-identity:%s:%s:identitypool/%s', region, account_id, id);


-- aws_directconnect_connections
ALTER TABLE aws_directconnect_connections ADD COLUMN arn TEXT;

UPDATE aws_directconnect_connections
SET arn = format('arn:aws:directconnect:%s:%s:dxcon/%s', region, account_id, id);


-- aws_directconnect_gateways
ALTER TABLE aws_directconnect_gateways ADD COLUMN arn TEXT;

UPDATE aws_directconnect_gateways
SET arn = format('arn:aws:directconnect::%s:dx-gateway/%s', account_id, id);


-- aws_directconnect_lags
ALTER TABLE aws_directconnect_lags ADD COLUMN arn TEXT;

UPDATE aws_directconnect_lags
SET arn = format('arn:aws:directconnect:%s:%s:dxlag/%s', region, account_id, id);


-- aws_directconnect_virtual_interfaces
ALTER TABLE aws_directconnect_virtual_interfaces ADD COLUMN arn TEXT;

UPDATE aws_directconnect_virtual_interfaces
SET arn = format('arn:aws:directconnect:%s:%s:dxvif/%s', region, account_id, id);


-- aws_ec2_flow_logs
ALTER TABLE aws_ec2_flow_logs ADD COLUMN arn TEXT;

UPDATE aws_ec2_flow_logs
SET arn = format('arn:aws:ec2:%s:%s:vpc-flow-log/%s', region, account_id, id);


-- aws_ec2_images
ALTER TABLE aws_ec2_images ADD COLUMN arn TEXT;

UPDATE aws_ec2_images
SET arn = format('arn:aws:ec2:%s::image/%s', region, id);


-- aws_ec2_instances
ALTER TABLE aws_ec2_instances ADD COLUMN arn TEXT;

UPDATE aws_ec2_instances
SET arn = format('arn:aws:ec2:%s:%s:instance/%s', region, account_id, id);

ALTER TABLE aws_ec2_instance_network_interfaces ADD COLUMN arn TEXT;

UPDATE aws_ec2_instance_network_interfaces t
SET arn = format('arn:aws:ec2:%s:%s:network-interface/%s', a.region, a.account_id, t.network_interface_id)
FROM aws_ec2_instances a
WHERE a.cq_id = t.instance_cq_id;


-- aws_ec2_internet_gateways
ALTER TABLE aws_ec2_internet_gateways ADD COLUMN arn TEXT;

UPDATE aws_ec2_internet_gateways
SET arn = format('arn:aws:ec2:%s:%s:internet-gateway/%s', region, account_id, id);


-- aws_ec2_nat_gateways
ALTER TABLE aws_ec2_nat_gateways ADD COLUMN arn TEXT;

UPDATE aws_ec2_nat_gateways
SET arn = format('arn:aws:ec2:%s:%s:natgateway/%s', region, account_id, id);


-- aws_ec2_network_acls
ALTER TABLE aws_ec2_network_acls ADD COLUMN arn TEXT;

UPDATE aws_ec2_network_acls
SET arn = format('arn:aws:ec2:%s:%s:network-acl/%s', region, account_id, id);


-- aws_ec2_route_tables
ALTER TABLE aws_ec2_route_tables ADD COLUMN arn TEXT;

UPDATE aws_ec2_route_tables
SET arn = format('arn:aws:ec2:%s:%s:route-table/%s', region, account_id, id);


-- aws_ec2_transit_gateways
ALTER TABLE aws_ec2_transit_gateways RENAME COLUMN transit_gateway_arn TO arn;


-- aws_ec2_vpc_endpoints
ALTER TABLE aws_ec2_vpc_endpoints ADD COLUMN arn TEXT;

UPDATE aws_ec2_vpc_endpoints
SET arn = format('arn:aws:ec2:%s:%s:vpc-endpoint/%s', region, account_id, id);


-- aws_ec2_vpc_peering_connections
ALTER TABLE aws_ec2_vpc_peering_connections ADD COLUMN arn TEXT;

UPDATE aws_ec2_vpc_peering_connections
SET arn = format('arn:aws:ec2:%s:%s:vpc-peering-connection/%s', region, account_id, id);


-- aws_ec2_vpcs
ALTER TABLE aws_ec2_vpcs ADD COLUMN arn TEXT;

UPDATE aws_ec2_vpcs
SET arn = format('arn:aws:ec2:%s:%s:vpc/%s', region, account_id, id);


-- aws_ec2_vpn_gateways
ALTER TABLE aws_ec2_vpn_gateways ADD COLUMN arn TEXT;

UPDATE aws_ec2_vpn_gateways
SET arn = format('arn:aws:ec2:%s:%s:vpn-gateway/%s', region, account_id, id);


-- aws_efs_filesystems
ALTER TABLE aws_efs_filesystems RENAME COLUMN file_system_arn TO arn;


-- aws_elbv1_load_balancers
ALTER TABLE aws_elbv1_load_balancers ADD COLUMN arn TEXT;

UPDATE aws_elbv1_load_balancers
SET arn = format('arn:aws:elasticloadbalancing:%s:%s:loadbalancer/%s', region, account_id, name);


-- aws_rds_clusters
ALTER TABLE aws_rds_clusters RENAME COLUMN db_cluster_arn TO arn;


-- aws_redshift_clusters
ALTER TABLE aws_redshift_clusters ADD COLUMN arn TEXT;

UPDATE aws_redshift_clusters
SET arn = format('arn:aws:redshift:%s:%s:cluster/%s', region, account_id, id);


-- aws_redshift_subnet_groups
ALTER TABLE aws_redshift_subnet_groups ADD COLUMN arn TEXT;

UPDATE aws_redshift_subnet_groups
SET arn = format('arn:aws:redshift:%s:%s:subnetgroup/%s', region, account_id, cluster_subnet_group_name);


-- aws_guardduty_detectors
ALTER TABLE aws_guardduty_detectors ADD COLUMN arn TEXT;

UPDATE aws_guardduty_detectors
SET arn = format('arn:aws:guardduty:%s:%s:detector/%s', region, account_id, id);
