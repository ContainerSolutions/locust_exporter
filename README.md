# Locust Exporter

Prometheus exporter for [Locust](https://github.com/locustio/locust). This exporter was inspired by [mbolek/locust_exporter](https://github.com/mbolek/locust_exporter).

[![CircleCI](https://circleci.com/gh/ContainerSolutions/locust_exporter.svg?style=svg)](https://circleci.com/gh/ContainerSolutions/locust_exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/containersol/locust_exporter.svg)](https://hub.docker.com/r/containersol/locust_exporter/tags)

## Quick Start

This package is available for Docker:

1. Run Locust Master

2. Run Locust Exporter

```bash
docker run --net=host containersol/locust_exporter
```

## Building and running

The default way to build is:

```bash
go get github.com/ContainerSolutions/locust_exporter
cd ${GOPATH-$HOME/go}/src/github.com/ContainerSolutions/locust_exporter/
go run main.go
```

### Flags

- `locust.uri`
  Address of Locust. Default is `http://localhost:8089`.

- `locust.timeout`
  Timeout request to Locust. Default is `5s`.

- `web.listen-address`
  Address to listen on for web interface and telemetry. Default is `:9646`.

- `web.telemetry-path`
  Path under which to expose metrics. Default is `/metrics`.

- `log.level`
  Set logging level: one of `debug`, `info`, `warn`, `error`, `fatal`

- `log.format`
  Set the log output target and format. e.g. `logger:syslog?appname=bob&local=7` or `logger:stdout?json=true`
  Defaults to `logger:stderr`.

### Environment Variables

The following environment variables configure the exporter:

- `LOCUST_EXPORTER_URI`
  Address of Locust. Default is `http://localhost:8089`.

- `LOCUST_EXPORTER_TIMEOUT`
  Timeout reqeust to Locust. Default is `5s`.

- `LOCUST_EXPORTER_WEB_LISTEN_ADDRESS`
  Address to listen on for web interface and telemetry. Default is `:9646`.

- `LOCUST_EXPORTER_WEB_TELEMETRY_PATH`
  Path under which to expose metrics. Default is `/metrics`.

### Screenshot

![locust exporter](locust_exporter.png)
