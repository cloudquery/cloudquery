# Table: aws_apigatewayv2_api_stages

This table shows data for Amazon API Gateway v2 API Stages.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-stages.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn (PK)|`utf8`|
|stage_name|`utf8`|
|access_log_settings|`json`|
|api_gateway_managed|`bool`|
|auto_deploy|`bool`|
|client_certificate_id|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|default_route_settings|`json`|
|deployment_id|`utf8`|
|description|`utf8`|
|last_deployment_status_message|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|route_settings|`json`|
|stage_variables|`json`|
|tags|`json`|

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


