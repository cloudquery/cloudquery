---
title: Reducing Alert Fatigue with the PagerDuty Plugin
tag: announcement
date: 2022/12/21
description: >-
  PagerDuty is a popular incident management tool used by many organizations. With the new PagerDuty plugin released today,
  we hope to help provide organizations insights on their incident management data, and help them reduce alert fatigue.

author: shimon
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

PagerDuty is a popular incident management tool used by many organizations. With the new PagerDuty plugin released today,
we hope to help provide organizations insights on their incident management data, and help them reduce alert fatigue.

Among other possible use-cases, you can find out how long it takes to acknowledge and resolve incidents in your organizations,
which users acknowledge and work on most incidents, and how many incidents each service generates. The example
queries below limit the data to the last 3 months, but you can of course change the time-frame, or write queries
that show month-by-month progress in the metrics.

## Getting Started with CloudQuery and PagerDuty

- Pick a destination to download your PagerDuty data into (this guide provides example queries for PostgreSQL). 
  You can find our list of supported destinations [here](/docs/plugins/destinations/overview).
- Create a [PagerDuty authorization token](https://support.pagerduty.com/docs/api-access-keys#section-generating-a-general-access-rest-api-key).
- Follow our [quickstart](/docs/quickstart) guide, and create a `pagerduty.yml` configuration file:
  
  ```yaml
  kind: source
  spec:
    name: "pagerduty"
    path: cloudquery/pagerduty
    version: "VERSION_SOURCE_PAGERDUTY" 
    destinations: ["postgresql"]
    tables: ["*"]
  ```

- Export your PagerDuty authorization token, and run `cloudquery`.
  ```bash
  PAGERDUTY_AUTH_TOKEN=<YOUR_AUTH_TOKEN>
  cloudquery sync .
  ```

## Use Cases
Here are a few SQL queries and ideas on insight you can gain from your PagerDuty data to help reduce alert fatigue.

### Average time to acknowledge and resolve incidents

The average time it takes to respond and resolve incidents can be a good indicator of how well your team is handling incidents.
- This query calculates the average time it took to respond to incidents in the last 3 months, grouped by priority:
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

- This query calculates the average time it takes to resolve incidents, grouped by priority:
  ```sql
  WITH incident_resolution_logs AS (
    SELECT pagerduty_incidents.id,
           pagerduty_incidents.priority->>'name' AS priority,
           pagerduty_incident_log_entries.created_at - pagerduty_incidents.created_at AS time_to_log
    FROM pagerduty_incidents 
    INNER JOIN pagerduty_incident_log_entries 
    ON pagerduty_incidents.id = pagerduty_incident_log_entries.incident->>'id'
    WHERE pagerduty_incident_log_entries.type = 'resolve_log_entry'
    AND pagerduty_incidents.created_at > NOW() - INTERVAL '3 months'
  ),
  incident_resolution_time AS ( -- Make sure only the last resolution is used (in case of multiple resolutions)
    SELECT id, 
           priority, 
           MAX(time_to_log) AS time_to_resolve
    FROM incident_resolution_logs
    GROUP BY id, priority
  )
  SELECT priority, AVG(time_to_resolve) AS average_time_to_resolve 
  FROM incident_resolution_time
  GROUP BY priority
  ```
  ```text
  priority | average_time_to_resolve
  ---------+-------------------------
  P1       | 08:12:37
  P2       | 1 day 04:22:24:04
  P3       | 3 days 01:02:00
  ```

### Find out which users acknowledged most incidents in the last 3 months

The following query finds out which members of your team take most of the load on working on incidents,
by finding out which users acknowledge the most incidents.

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

### Top 10 services that generate most incidents

This query finds out which services generated the most incidents in the last 3 months.

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