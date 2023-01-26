# PagerDuty Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", "pagerduty")}/>

The CloudQuery PagerDuty plugin extracts PagerDuty resources. It is based on [The PagerDuty Go SDK](https://github.com/PagerDuty/go-pagerduty) and the [PagerDuty REST API](https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUw-rest-api-v2-overview).

## Authentication

In order to authenticate with your PagerDuty account, you will need a [PagerDuty authorization token](https://support.pagerduty.com/docs/api-access-keys#section-generating-a-general-access-rest-api-key).
CloudQuery supports two methods of reading the authorization token:
- From a `~/.pd.yml` file, such as:
  ```yaml
  authtoken: <YOUR_AUTH_TOKEN>
  ```
- From an environment variable `PAGERDUTY_AUTH_TOKEN`.

## Example Queries

The following example queries are SQL queries in the PostgreSQL flavour (i.e. can be used with the 
[PostgreSQL destination](/docs/plugins/destinations/postgresql/overview)).

### Top 10 services that generate the most incidents (last 3 months)

```sql
  SELECT pagerduty_services.id AS service_id,
         pagerduty_services.name AS service_name,
         COUNT(pagerduty_incidents.id) AS incident_count
  FROM pagerduty_services
  INNER JOIN pagerduty_incidents
  ON pagerduty_services.id = pagerduty_incidents.service->>'id'
  WHERE pagerduty_incidents.created_at > NOW() - INTERVAL '3 months'
  GROUP BY service_id, service_name
  ORDER BY incident_count DESC
  LIMIT 10
```

```text
 service_id |  service_name   | incident_count
------------+-----------------+----------------
 PYS6MP5    | UnstableService | 25   
 PAZ9U1C    | StableService   | 3
```

### Average time to respond to a query, grouped by priority (last 3 months)

```sql
WITH incident_ack_logs AS (
  SELECT pagerduty_incidents.id AS incident_id,
         pagerduty_incidents.priority->>'name' AS priority,
         pagerduty_incident_log_entries.created_at - pagerduty_incidents.created_at AS time_to_log
  FROM pagerduty_incidents 
  INNER JOIN pagerduty_incident_log_entries 
  ON pagerduty_incidents.id = pagerduty_incident_log_entries.incident->>'id'
  WHERE pagerduty_incident_log_entries.type = 'acknowledge_log_entry'
  AND pagerduty_incidents.created_at > NOW() - INTERVAL '3 months'
),
incident_ack_time AS ( -- Make sure only the first acknowledgement is used (incidents may be acknowledged twice)
  SELECT incident_id, 
         priority, 
         MIN(time_to_log) AS time_to_ack
  FROM incident_ack_logs
  GROUP BY incident_id, priority
)
SELECT priority, AVG(time_to_ack) AS average_time_to_ack
FROM incident_ack_time
GROUP BY priority
```

```text
priority | average_time_to_ack
----------+---------------------
P1       | 00:05:10
P2       | 00:30:00
P3       | 06:20:00
```

### Which users acknowledged the most incidents (last 3 months)

```sql
  WITH incident_ack_logs AS (
    SELECT pagerduty_incidents.id AS incident_id,
           pagerduty_incident_log_entries.agent->>'id' AS agent_id
    FROM pagerduty_incidents
    INNER JOIN pagerduty_incident_log_entries 
    ON pagerduty_incidents.id = pagerduty_incident_log_entries.incident->>'id'
    WHERE pagerduty_incident_log_entries.type = 'acknowledge_log_entry'
    AND pagerduty_incidents.created_at > NOW() - INTERVAL '3 months'
  )
  SELECT pagerduty_users.id AS user_id, 
         pagerduty_users.name AS user_name, 
         COUNT(DISTINCT incident_id) AS acknowledge_count FROM
  pagerduty_users INNER JOIN incident_ack_logs
  ON pagerduty_users.id = incident_ack_logs.agent_id
  GROUP BY user_id, user_name
  ORDER BY acknowledge_count DESC
```

```text
 user_id |  user_name  | acknowledge_count
---------+-------------+-------------------
 PDYR2Y8 | John        | 15
 PDYR2Y9 | Dave        | 12
 PDYR2Z5 | Jane        | 8
```