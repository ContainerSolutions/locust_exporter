# Locust Exporter

The Locust Exporter for Prometheus based on [mbolek/locust_exporter](https://github.com/mbolek/locust_exporter).

[![Build Status](https://travis-ci.org/ContainerSolutions/locust_exporter.svg?branch=master)](https://travis-ci.org/ContainerSolutions/locust_exporter)

This is a simple exporter for http://locust.io metrics. You get all the necessary details about current tests and the state of the locust.

## Quick Start

This package is available as a container:

``` bash
docker run --net=host -e LOCUST="localhost:8089" containersol/locust_exporter
```

### Environment Variables

The following environment variables configure the exporter:

* `LOCUST`
  Address of Locust. Default is `localhost:8089`.

* `LISTENER_PORT`
  Port of metrics. Default is `9646`.

## Info

Errors and requests stats are added with the method and path labels - **BE CAREFUL** - if you have a lot of endpoints. It is probably better to group the endpoints in your locustfile: [see here](http://docs.locust.io/en/latest/writing-a-locustfile.html#grouping-requests-to-urls-with-dynamic-parameters).

## Requirements

Requirements:

* prometheus_client  
