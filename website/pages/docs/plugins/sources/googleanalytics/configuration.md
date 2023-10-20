# Google Analytics Source Plugin Configuration Reference

## Example

This example syncs from Google Analytics to a Postgres destination.
The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

## Google Analytics Spec

This is the (nested) spec used by the Google Analytics source plugin:

- `property_id` (`string`) (required):

  A Google Analytics GA4 [property](https://support.google.com/analytics/answer/9304153#property) identifier whose events are tracked.
  To learn more, see where to [find your Property ID](https://developers.google.com/analytics/devguides/reporting/data/v1/property-id).

  Supported formats:

  - A plain property ID (example: `1234`)

  - Prefixed with `properties/` (example: `properties/1234`)

- `reports` ([`[]report`](#google-analytics-report-spec)) (required):

  Reports to be fetched from Google Analytics.

- `start_date` (`string`) (optional) (default: date 7 days prior to the sync start):

  A date in `YYYY-MM-DD` format (example: `2023-05-15`).
  If not specified, the start date will be the one that is 7 days prior to the sync start date.

- `oauth` ([OAuth spec](#google-analytics-oauth-spec)) (optional) (default: empty)

  OAuth spec for authorization in Google Analytics.

- `concurrency` (`integer`) (optional) (default: `10000`):

  The best effort maximum number of Go routines to use.
  Lower this number to reduce memory usage.

### Google Analytics OAuth spec

OAuth spec to authenticate with Google Analytics.
[Google Analytics Data API v1](https://developers.google.com/analytics/devguides/reporting/data/v1)
requires OAuth authorization for `https://www.googleapis.com/auth/analytics.readonly` scope to run reports.

- `access_token` (`string`) (optional) (default: `""`)

  An access token that you generated authorizing for `https://www.googleapis.com/auth/analytics.readonly` scope
  (e.g., by using [OAuth 2.0 Playground](https://developers.google.com/oauthplayground/)).

- `client_id` (`string`) (optional) (default: `""`)

  OAuth 2.0 Client ID.
  Required if `access_token` is empty.

- `client_secret` (`string`) (optional) (default: `""`)

  OAuth 2.0 Client secret.
  Required if `access_token` is empty.

### Google Analytics Report spec

Report specification will be transformed into a Google Analytics Data API v1
[report](https://developers.google.com/analytics/devguides/reporting/data/v1/basics#reports).
The option structure follows:

- `name` (`string`) (required):

  Name of the report.
  It will be translated into a table name as `ga_` prefix followed by report name in snake case.

- `dimensions` (`[]string`) (optional) (default: empty)

  A list of Google Analytics Data API v1 [dimensions](https://developers.google.com/analytics/devguides/reporting/data/v1/api-schema#dimensions).
  At most `9` dimensions can be specified per report.

- `metrics` ([`[]metric`](#google-analytics-metric-spec)) (required)

  A list of Google Analytics Data API v1 [metrics](https://developers.google.com/analytics/devguides/reporting/data/v1/api-schema#metrics).
  Expressions are supported, too.

- `keep_empty_rows` (`boolean`) (optional) (default: `false`)

  Whether empty rows should be captured, too.

#### Google Analytics metric spec

Metric spec that is based on Google Analytics Data API v1
[Metric](https://developers.google.com/analytics/devguides/reporting/data/v1/rest/v1beta/Metric) parameter.

- `name` (`string`) (required)

  A name or alias (if `expression` is specified) of the requested metric.

- `expression` (`string`) (optional) (default: `""`)

  A mathematical expression for derived metrics.

- `invisible` (`boolean`) (optional) (default: `false`)

  Indicates if a metric is invisible in the report response.
  This allows creating more complex requests, while also not saving the intermediate results.
