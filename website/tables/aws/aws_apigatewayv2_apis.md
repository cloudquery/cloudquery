# Table: aws_apigatewayv2_apis

This table shows data for Amazon API Gateway v2 APIs.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_apigatewayv2_apis:
  - [aws_apigatewayv2_api_authorizers](aws_apigatewayv2_api_authorizers)
  - [aws_apigatewayv2_api_deployments](aws_apigatewayv2_api_deployments)
  - [aws_apigatewayv2_api_integrations](aws_apigatewayv2_api_integrations)
  - [aws_apigatewayv2_api_models](aws_apigatewayv2_api_models)
  - [aws_apigatewayv2_api_routes](aws_apigatewayv2_api_routes)
  - [aws_apigatewayv2_api_stages](aws_apigatewayv2_api_stages)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|
|protocol_type|`utf8`|
|route_selection_expression|`utf8`|
|api_endpoint|`utf8`|
|api_gateway_managed|`bool`|
|api_id|`utf8`|
|api_key_selection_expression|`utf8`|
|cors_configuration|`json`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|disable_execute_api_endpoint|`bool`|
|disable_schema_validation|`bool`|
|import_info|`list<item: utf8, nullable>`|
|tags|`json`|
|version|`utf8`|
|warnings|`list<item: utf8, nullable>`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### API Gateway REST and WebSocket API logging should be enabled

```sql
(
  SELECT
    DISTINCT
    'API Gateway REST and WebSocket API logging should be enabled' AS title,
    r.account_id,
    'arn:' || 'aws' || ':apigateway:' || r.region || ':/restapis/' || r.id
      AS resource_id,
    CASE
    WHEN s.logging_level NOT IN ('"ERROR"', '"INFO"') THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    view_aws_apigateway_method_settings AS s
    LEFT JOIN aws_apigateway_rest_apis AS r ON s.rest_api_arn = r.arn
)
UNION
  (
    SELECT
      DISTINCT
      'API Gateway REST and WebSocket API logging should be enabled' AS title,
      a.account_id,
      'arn:' || 'aws' || ':apigateway:' || a.region || ':/apis/' || a.id
        AS resource_id,
      CASE
      WHEN s.default_route_settings->>'LoggingLevel' IN (NULL, 'OFF')
      THEN 'fail'
      ELSE 'pass'
      END
        AS status
    FROM
      aws_apigatewayv2_api_stages AS s
      LEFT JOIN aws_apigatewayv2_apis AS a ON s.api_arn = a.arn
  );
```

### Find all API Gateway V2 instances (HTTP and Webhook) that are publicly accessible

```sql
SELECT
  'Find all API Gateway V2 instances (HTTP and Webhook) that are publicly accessible'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_apigatewayv2_apis;
```


