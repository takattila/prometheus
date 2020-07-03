# Example Sevice - Grafana Dashboard

[Back to README.md](./README.md#table-of-contents)

```json
{
  "__inputs": [
    {
      "name": "DS_PROMETHEUS",
      "label": "Prometheus",
      "description": "",
      "type": "datasource",
      "pluginId": "prometheus",
      "pluginName": "Prometheus"
    }
  ],
  "__requires": [
    {
      "type": "panel",
      "id": "gauge",
      "name": "Gauge",
      "version": ""
    },
    {
      "type": "grafana",
      "id": "grafana",
      "name": "Grafana",
      "version": "6.1.6"
    },
    {
      "type": "panel",
      "id": "grafana-piechart-panel",
      "name": "Pie Chart",
      "version": "1.5.0"
    },
    {
      "type": "panel",
      "id": "graph",
      "name": "Graph",
      "version": ""
    },
    {
      "type": "datasource",
      "id": "prometheus",
      "name": "Prometheus",
      "version": "1.0.0"
    },
    {
      "type": "panel",
      "id": "singlestat",
      "name": "Singlestat",
      "version": ""
    }
  ],
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "id": null,
  "links": [],
  "panels": [
    {
      "gridPos": {
        "h": 5,
        "w": 8,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "links": [],
      "options": {
        "maxValue": 100,
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "light-green",
            "index": 0,
            "value": null
          },
          {
            "color": "dark-green",
            "index": 1,
            "value": 80
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "mean",
          "suffix": "",
          "unit": "none"
        }
      },
      "pluginVersion": "6.1.6",
      "targets": [
        {
          "expr": "cpu_usage{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}",
          "format": "time_series",
          "instant": true,
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "Handler 1 Gauge",
      "type": "gauge"
    },
    {
      "datasource": "${DS_PROMETHEUS}",
      "gridPos": {
        "h": 5,
        "w": 8,
        "x": 8,
        "y": 0
      },
      "id": 12,
      "links": [],
      "options": {
        "maxValue": "20",
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "light-orange",
            "index": 0,
            "value": null
          },
          {
            "color": "dark-orange",
            "index": 1,
            "value": 15
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "last",
          "suffix": "",
          "unit": "none"
        }
      },
      "pluginVersion": "6.1.6",
      "targets": [
        {
          "expr": "stat_goroutines:count{app=\"exampleService\",env=\"test\"} ",
          "format": "time_series",
          "instant": true,
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "Service Goroutines",
      "type": "gauge"
    },
    {
      "gridPos": {
        "h": 5,
        "w": 8,
        "x": 16,
        "y": 0
      },
      "id": 4,
      "links": [],
      "options": {
        "maxValue": 100,
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "light-blue",
            "index": 0,
            "value": null
          },
          {
            "color": "dark-blue",
            "index": 1,
            "value": 80
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "mean",
          "suffix": "",
          "unit": "none"
        }
      },
      "pluginVersion": "6.1.6",
      "targets": [
        {
          "expr": "cpu_usage{app=\"exampleService\",env=\"test\",handler=\"MyHandler2\"}",
          "format": "time_series",
          "instant": true,
          "interval": "",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "Handler 2 Gauge",
      "type": "gauge"
    },
    {
      "gridPos": {
        "h": 5,
        "w": 8,
        "x": 0,
        "y": 5
      },
      "id": 60,
      "links": [],
      "options": {
        "maxValue": 100,
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "super-light-red",
            "index": 0,
            "value": null
          },
          {
            "color": "dark-red",
            "index": 1,
            "value": 80
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "mean",
          "suffix": "",
          "unit": "percent"
        }
      },
      "pluginVersion": "6.1.6",
      "targets": [
        {
          "expr": "stat_cpu_usage:percent{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "instant": true,
          "intervalFactor": 1,
          "legendFormat": "",
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "CPU Usage",
      "type": "gauge"
    },
    {
      "aliasColors": {},
      "bars": true,
      "cacheTimeout": null,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 5,
        "w": 16,
        "x": 8,
        "y": 5
      },
      "id": 61,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": false,
        "total": false,
        "values": false
      },
      "lines": false,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pluginVersion": "6.1.6",
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "CPU usage",
          "color": "#FF7383"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "stat_cpu_usage:percent{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "instant": false,
          "intervalFactor": 1,
          "legendFormat": "CPU usage",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "CPU Usage",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#56A64B",
        "#299c46"
      ],
      "format": "percent",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 12,
        "x": 0,
        "y": 10
      },
      "id": 28,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#96D98D",
        "full": false,
        "lineColor": "#56A64B",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:used_percent{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "instant": false,
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "vm total",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Used memory (System)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPrefix": false,
      "colorValue": true,
      "colors": [
        "#299c46",
        "#3274D9",
        "#d44a3a"
      ],
      "format": "percent",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 12,
        "x": 12,
        "y": 10
      },
      "id": 30,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#8AB8FF",
        "full": false,
        "lineColor": "#3274D9",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "100 - stat_memory_usage:used_percent{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "intervalFactor": 1,
          "legendFormat": "free memory",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Free memory (System)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#37872D",
        "#299c46"
      ],
      "description": "Sys is the total bytes of memory obtained from the OS",
      "format": "decbytes",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 6,
        "x": 0,
        "y": 13
      },
      "id": 32,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#96D98D",
        "full": false,
        "lineColor": "#56A64B",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:sys{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "app sys",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Memory: sys (Application)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#37872D",
        "#299c46"
      ],
      "description": "Alloc is bytes of allocated heap objects",
      "format": "decbytes",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 6,
        "x": 6,
        "y": 13
      },
      "id": 34,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#96D98D",
        "full": false,
        "lineColor": "#56A64B",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:alloc{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "instant": false,
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "app alloc",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Memory: alloc (Application)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#3274D9",
        "#299c46"
      ],
      "description": " HeapSys is bytes of heap memory obtained from the OS",
      "format": "decbytes",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 6,
        "x": 12,
        "y": 13
      },
      "id": 36,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#8AB8FF",
        "full": false,
        "lineColor": "#3274D9",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:heapsys{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "app heapsys",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Memory: heapsys (Application)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#3274D9",
        "#299c46"
      ],
      "description": "HeapInuse is bytes in in-use spans",
      "format": "decbytes",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 6,
        "x": 18,
        "y": 13
      },
      "id": 38,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#8AB8FF",
        "full": false,
        "lineColor": "#3274D9",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:heapinuse{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "app heapinuse",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Memory: heapinuse (Application)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#37872D",
        "#299c46"
      ],
      "description": "Frees is the cumulative count of heap objects freed",
      "format": "decbytes",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 12,
        "x": 0,
        "y": 16
      },
      "id": 40,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#96D98D",
        "full": false,
        "lineColor": "#56A64B",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:frees{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "app frees",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Memory: frees (Application)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorPostfix": false,
      "colorPrefix": true,
      "colorValue": true,
      "colors": [
        "#d44a3a",
        "#3274D9",
        "#299c46"
      ],
      "description": " NumGC is the number of completed GC cycles",
      "format": "none",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 3,
        "w": 12,
        "x": 12,
        "y": 16
      },
      "id": 42,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "#8AB8FF",
        "full": false,
        "lineColor": "#3274D9",
        "show": true
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "stat_memory_usage:numgc{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "app numgc",
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "Memory: numgc (Application)",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 19
      },
      "id": 58,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "elapsed time",
          "color": "#3274D9"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    rate(\n      elapsed_time_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler2\"}[5m]\n    ) \n) ",
          "format": "time_series",
          "interval": "1s",
          "intervalFactor": 1,
          "legendFormat": "elapsed time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "elapsed_time",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "s",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 27
      },
      "id": 52,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    avg(increase(\n        execution_time_nano_sec_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[1m]\n    )) by (le)\n) ",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "execution time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "execution_time_nano_sec",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "ns",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 27
      },
      "id": 50,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    avg(increase(\n        execution_time_micro_sec_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[1m]\n    )) by (le)\n) ",
          "format": "time_series",
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "execution time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "execution_time_micro_sec",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "µs",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 35
      },
      "id": 48,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "execution time",
          "color": "#8F3BB8"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    rate(\n        execution_time_milli_sec_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[5m]\n    )\n) ",
          "format": "time_series",
          "interval": "10ms",
          "intervalFactor": 1,
          "legendFormat": "execution time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "execution_time_milli_sec",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "ms",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 43
      },
      "id": 54,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "execution time",
          "color": "#FF780A"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    rate(\n      execution_time_seconds_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[5m]\n    ) \n) ",
          "format": "time_series",
          "interval": "1s",
          "intervalFactor": 1,
          "legendFormat": "execution time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "execution_time_seconds",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "s",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 51
      },
      "id": 56,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    rate(\n      execution_time_minutes_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[5m]\n    ) \n) ",
          "format": "time_series",
          "interval": "1s",
          "intervalFactor": 1,
          "legendFormat": "execution time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "execution_time_minutes",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "m",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 59
      },
      "id": 6,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    rate(\n        response_time_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[1m]\n    )\n)",
          "format": "time_series",
          "instant": false,
          "interval": "10ms",
          "intervalFactor": 1,
          "legendFormat": "{{le}}",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "Handler 1 Response Time (AVG)",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "s",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "datasource": "${DS_PROMETHEUS}",
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 59
      },
      "id": 8,
      "interval": "1s",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "response time",
          "color": "#3274D9"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.9,\n    rate(\n        response_time_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler2\"}[5m]\n    )\n) ",
          "format": "time_series",
          "interval": "10ms",
          "intervalFactor": 1,
          "legendFormat": "response time",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "Handler 2 Response Time (AVG)",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "s",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 9,
        "w": 8,
        "x": 0,
        "y": 67
      },
      "id": 10,
      "interval": "",
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "Handler 2",
          "color": "#1250B0"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "histogram_quantile(0.5,\n    avg(increase(\n        response_time_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[5m]\n    )) by (le)\n) ",
          "format": "time_series",
          "interval": "10ms",
          "intervalFactor": 1,
          "legendFormat": "Handler 1",
          "refId": "A"
        },
        {
          "expr": "histogram_quantile(0.5,\n    avg(increase(\n        response_time_bucket{app=\"exampleService\",env=\"test\",handler=\"MyHandler2\"}[5m]\n    )) by (le)\n) ",
          "format": "time_series",
          "interval": "10ms",
          "intervalFactor": 1,
          "legendFormat": "Handler 2",
          "refId": "B"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "Response Times of Handlers (AVG / Comparison)",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "s",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 9,
        "w": 8,
        "x": 8,
        "y": 67
      },
      "id": 20,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "stat_memory_usage:total{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "intervalFactor": 1,
          "legendFormat": "vm total",
          "refId": "A"
        },
        {
          "expr": "stat_memory_usage:avail{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "intervalFactor": 1,
          "legendFormat": "vm avail",
          "refId": "B"
        },
        {
          "expr": "stat_memory_usage:used{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "intervalFactor": 1,
          "legendFormat": "vm used",
          "refId": "C"
        },
        {
          "expr": "stat_memory_usage:free{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "intervalFactor": 1,
          "legendFormat": "vm free",
          "refId": "D"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "Memory Usage (System)",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "decbytes",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 9,
        "w": 8,
        "x": 16,
        "y": 67
      },
      "id": 24,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [
        {
          "alias": "cpu usage",
          "color": "#E02F44"
        }
      ],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "stat_cpu_usage:percent{app=\"exampleService\",env=\"test\"}",
          "format": "time_series",
          "intervalFactor": 1,
          "legendFormat": "cpu usage",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "CPU Usage",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "percent",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "breakPoint": "50%",
      "cacheTimeout": null,
      "combine": {
        "label": "Others",
        "threshold": 0
      },
      "fontSize": "80%",
      "format": "short",
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 0,
        "y": 76
      },
      "id": 62,
      "interval": null,
      "legend": {
        "show": true,
        "values": true
      },
      "legendType": "Right side",
      "links": [],
      "maxDataPoints": 1,
      "nullPointMode": "connected",
      "pieType": "donut",
      "strokeWidth": 1,
      "targets": [
        {
          "expr": "response_status{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}\n",
          "format": "time_series",
          "instant": true,
          "intervalFactor": 1,
          "legendFormat": "{{statuscode}}",
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "Pie Chart - Status codes of MyHandler1",
      "type": "grafana-piechart-panel",
      "valueName": "current"
    },
    {
      "aliasColors": {},
      "bars": true,
      "cacheTimeout": null,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 76
      },
      "id": 63,
      "legend": {
        "alignAsTable": false,
        "avg": false,
        "current": false,
        "hideEmpty": false,
        "max": true,
        "min": true,
        "rightSide": true,
        "show": true,
        "total": false,
        "values": true
      },
      "lines": false,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "sum(\n  increase(\n    response_status{app=\"exampleService\",env=\"test\",handler=\"MyHandler1\"}[5m]\n  )\n) by(statuscode)\n",
          "format": "time_series",
          "instant": false,
          "intervalFactor": 1,
          "legendFormat": "{{statuscode}}",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "Timeline - Status codes of MyHandler1",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1024,
          "max": null,
          "min": null,
          "show": false
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "breakPoint": "50%",
      "cacheTimeout": null,
      "combine": {
        "label": "Others",
        "threshold": 0
      },
      "fontSize": "80%",
      "format": "short",
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 0,
        "y": 83
      },
      "id": 46,
      "interval": null,
      "legend": {
        "show": true,
        "values": true
      },
      "legendType": "Right side",
      "links": [],
      "maxDataPoints": 1,
      "nullPointMode": "connected",
      "pieType": "donut",
      "strokeWidth": 1,
      "targets": [
        {
          "expr": "response_status{app=\"exampleService\",env=\"test\",handler=\"MyHandler2\"}\n",
          "format": "time_series",
          "instant": true,
          "intervalFactor": 1,
          "legendFormat": "{{statuscode}}",
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "Pie Chart - Status codes of MyHandler2",
      "type": "grafana-piechart-panel",
      "valueName": "current"
    },
    {
      "aliasColors": {},
      "bars": true,
      "cacheTimeout": null,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 83
      },
      "id": 64,
      "legend": {
        "alignAsTable": false,
        "avg": false,
        "current": false,
        "hideEmpty": false,
        "max": true,
        "min": true,
        "rightSide": true,
        "show": true,
        "total": false,
        "values": true
      },
      "lines": false,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "sum(\n  increase(\n    response_status{app=\"exampleService\",env=\"test\",handler=\"MyHandler2\"}[5m]\n  )\n) by(statuscode)\n",
          "format": "time_series",
          "instant": false,
          "intervalFactor": 1,
          "legendFormat": "{{statuscode}}",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "Timeline - Status codes of MyHandler2",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1024,
          "max": null,
          "min": null,
          "show": false
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    }
  ],
  "refresh": "5s",
  "schemaVersion": 18,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-5m",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m",
      "15m",
      "30m",
      "1h",
      "2h",
      "1d"
    ],
    "time_options": [
      "5m",
      "15m",
      "1h",
      "6h",
      "12h",
      "24h",
      "2d",
      "7d",
      "30d"
    ]
  },
  "timezone": "",
  "title": "Example service",
  "uid": "JTnT-NWMk",
  "version": 79
}
```
[Back to README.md](./README.md#table-of-contents)