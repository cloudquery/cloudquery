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
