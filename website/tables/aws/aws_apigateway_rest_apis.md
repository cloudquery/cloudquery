# Table: aws_apigateway_rest_apis

This table shows data for Amazon API Gateway Rest APIs.

https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_apigateway_rest_apis:
  - [aws_apigateway_rest_api_authorizers](aws_apigateway_rest_api_authorizers)
  - [aws_apigateway_rest_api_deployments](aws_apigateway_rest_api_deployments)
  - [aws_apigateway_rest_api_documentation_parts](aws_apigateway_rest_api_documentation_parts)
  - [aws_apigateway_rest_api_documentation_versions](aws_apigateway_rest_api_documentation_versions)
  - [aws_apigateway_rest_api_gateway_responses](aws_apigateway_rest_api_gateway_responses)
  - [aws_apigateway_rest_api_models](aws_apigateway_rest_api_models)
  - [aws_apigateway_rest_api_request_validators](aws_apigateway_rest_api_request_validators)
  - [aws_apigateway_rest_api_resources](aws_apigateway_rest_api_resources)
  - [aws_apigateway_rest_api_stages](aws_apigateway_rest_api_stages)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|api_key_source|`utf8`|
|binary_media_types|`list<item: utf8, nullable>`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|disable_execute_api_endpoint|`bool`|
|endpoint_configuration|`json`|
|id|`utf8`|
|minimum_compression_size|`int64`|
|name|`utf8`|
|policy|`utf8`|
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

### Find all API Gateway instances that are publicly accessible

```sql
SELECT
  'Find all API Gateway instances that are publicly accessible' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN NOT ('{PRIVATE}' = t) THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_apigateway_rest_apis,
  jsonb_array_elements_text(endpoint_configuration->'Types') AS t;
```


