# Observability

In order to understand how CloudQuery is performing we added some basic observability capabilities.
The main way to do this currently is to use the CloudQuery log output.

The [CloudQuery Helm Chart](https://github.com/cloudquery/helm-charts/tree/main/charts/cloudquery) provides the ability to install Promtail into the cluster.
[Promtail](https://grafana.com/docs/loki/latest/clients/promtail/) is an agent which ships logs to a [Grafana Loki](https://grafana.com/oss/loki/) instance.
We decided on using Loki as the log sink as it is a simple, easy to use, and reliable solution that works across different Cloud providers.

## Install Promtail with CloudQuery Helm Chart

Add the following to your `values.yaml` file and make sure to replace the URL with the one of your Loki instance:

```yaml
promtail:
  enabled: true
  config:
    clients:
    - "http://loki-gateway/loki/api/v1/push"
```

The same applies for the Installation with our Terraform modules.

## Import the CloudQuery Overview Dashboard

<details>
<summary>Grafana Dashboard JSON Model</summary>

```json
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "grafana",
          "uid": "-- Grafana --"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "id": 12,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "datasource": {
        "type": "loki",
        "uid": "grafanacloud-logs"
      },
      "description": "",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": 1000,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "min": 0,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 9,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "9.0.6",
      "targets": [
        {
          "datasource": {
            "type": "loki",
            "uid": "grafanacloud-logs"
          },
          "editorMode": "builder",
          "expr": "sum by(table) (rate({app=\"cloudquery\"} |= `fetched successfully` | json | __error__=`` | message=`fetched successfully` [$__interval]))",
          "legendFormat": "{{table}}",
          "queryType": "range",
          "refId": "FetchedSuccessfully"
        }
      ],
      "title": "FetchedSuccessfully",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "loki",
        "uid": "grafanacloud-logs"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "always",
            "spanNulls": 1000,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "min": 0,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": [
          {
            "__systemRef": "hideSeriesFrom",
            "matcher": {
              "id": "byNames",
              "options": {
                "mode": "exclude",
                "names": [
                  "Errors"
                ],
                "prefix": "All except:",
                "readOnly": true
              }
            },
            "properties": [
              {
                "id": "custom.hideFrom",
                "value": {
                  "legend": false,
                  "tooltip": false,
                  "viz": true
                }
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 9,
        "w": 12,
        "x": 12,
        "y": 0
      },
      "id": 4,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "loki",
            "uid": "grafanacloud-logs"
          },
          "editorMode": "builder",
          "expr": "count(rate({app=\"cloudquery\"} |= `error` | json | __error__=`` | level=`error` [$__interval]))",
          "hide": false,
          "legendFormat": "Errors",
          "queryType": "range",
          "refId": "Errors"
        },
        {
          "datasource": {
            "type": "loki",
            "uid": "grafanacloud-logs"
          },
          "editorMode": "builder",
          "expr": "count(rate({app=\"cloudquery\"} |= `warn` | json | __error__=`` | level=`warn` [$__interval]))",
          "legendFormat": "Warnings",
          "queryType": "range",
          "refId": "Warnings"
        }
      ],
      "thresholds": [
        {
          "colorMode": "critical",
          "op": "gt",
          "value": 0,
          "visible": true
        }
      ],
      "title": "Exceptions",
      "type": "timeseries"
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 9
      },
      "id": 8,
      "panels": [],
      "title": "Errors",
      "type": "row"
    },
    {
      "datasource": {
        "type": "loki",
        "uid": "grafanacloud-logs"
      },
      "fieldConfig": {
        "defaults": {
          "custom": {
            "align": "auto",
            "displayMode": "auto",
            "filterable": false,
            "inspect": false
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": [
          {
            "matcher": {
              "id": "byName",
              "options": "Time"
            },
            "properties": [
              {
                "id": "custom.width",
                "value": 158
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Plugin"
            },
            "properties": [
              {
                "id": "custom.width",
                "value": 97
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Table"
            },
            "properties": [
              {
                "id": "custom.width",
                "value": 201
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Error"
            },
            "properties": [
              {
                "id": "custom.width",
                "value": 697
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 11,
        "w": 24,
        "x": 0,
        "y": 10
      },
      "id": 6,
      "options": {
        "footer": {
          "enablePagination": true,
          "fields": "",
          "reducer": [
            "sum"
          ],
          "show": false
        },
        "showHeader": true,
        "sortBy": []
      },
      "pluginVersion": "9.0.5",
      "targets": [
        {
          "datasource": {
            "type": "loki",
            "uid": "grafanacloud-logs"
          },
          "editorMode": "builder",
          "expr": "{app=\"cloudquery\"} |= `error` | json | __error__=`` | level=`error`",
          "legendFormat": "",
          "queryType": "range",
          "refId": "Errors"
        }
      ],
      "title": "Errors",
      "transformations": [
        {
          "id": "extractFields",
          "options": {
            "replace": true,
            "source": "Line"
          }
        },
        {
          "id": "organize",
          "options": {
            "excludeByName": {
              "Region": true,
              "account_id": true,
              "client_id": true,
              "details": true,
              "instance_id": true,
              "level": true,
              "resource": true,
              "resource_keys": true,
              "time": false,
              "timestamp": true,
              "type": true
            },
            "indexByName": {
              "@module": 1,
              "Region": 7,
              "account_id": 8,
              "client_id": 9,
              "error": 3,
              "instance_id": 6,
              "level": 5,
              "message": 4,
              "table": 2,
              "time": 0,
              "timestamp": 10
            },
            "renameByName": {
              "@module": "Plugin",
              "error": "Error",
              "message": "Message",
              "table": "Table",
              "time": "Time"
            }
          }
        }
      ],
      "type": "table"
    }
  ],
  "refresh": false,
  "schemaVersion": 36,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "Overview",
  "uid": "jFgNVxz4z",
  "version": 11,
  "weekStart": ""
}
```

</details>

## Using the Dashboard

The Overview Dashboard is a simple way to see how many fetches are running and how many warnings and errors are happening.
Alerts can easily be setup for warnings and errors.

![observability](/images/observability.png)
