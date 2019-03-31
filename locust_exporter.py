#!/usr/bin/python
# -*- coding: UTF-8 -*-
# Requires prometheus_client library:
# sudo pip install prometheus_client
from prometheus_client import start_http_server, Metric, REGISTRY
import json
import requests
import sys
import time
import os

class LocustCollector(object):
  def __init__(self, ep):
    self._ep = ep

  def collect(self):
    # Fetch the JSON
    url = 'http://' + self._ep + '/stats/requests'
    try:
        response = requests.get(url).content.decode('Utf-8')
    except requests.exceptions.ConnectionError:
        print("Failed to connect to Locust:", url)
        return None

    response = json.loads(response)

    stats_metrics=['avg_content_length','avg_response_time','current_rps','max_response_time','median_response_time','min_response_time','num_failures','num_requests',
                   'max_connect','max_dns_lookup','max_name_lookup','max_pre_transfer','max_server_processing','max_start_transfer','max_tcp_connection','max_tls_handshake','min_connect','min_dns_lookup','min_name_lookup','min_pre_transfer','min_server_processing','min_start_transfer','min_tcp_connection','min_tls_handshake']

    metric = Metric('locust_user_count', 'Swarmed users', 'gauge')
    metric.add_sample('locust_user_count', value=response['user_count'], labels={})
    yield metric

    metric = Metric('locust_response_time_95', 'Response Time 95th Percentile', 'gauge')
    metric.add_sample('locust_response_time_95', value=response["current_response_time_percentile_95"], labels={})
    yield metric
    if 'current_response_time_percentile_99' in response:
        metric = Metric('locust_response_time_99', 'Response Time 99th Percentile', 'gauge')
        metric.add_sample('locust_response_time_99', value=response["current_response_time_percentile_99"], labels={})
        yield metric

    metric = Metric('locust_errors', 'Locust requests errors', 'gauge')
    for err in response['errors']:
        metric.add_sample('locust_errors', value=err['occurences'], labels={'path':err['name'], 'method':err['method'], 'error': err['error']})
    yield metric

    if 'slave_count' in response:
        metric = Metric('locust_slave_count', 'Locust number of slaves', 'gauge')
        metric.add_sample('locust_slave_count', value=response['slave_count'], labels={})
        yield metric

    metric = Metric('locust_fail_ratio', 'Locust failure ratio', 'gauge')
    metric.add_sample('locust_fail_ratio', value=response['fail_ratio'], labels={})
    yield metric

    metric = Metric('locust_state', 'State of the locust swarm', 'gauge')
    metric.add_sample('locust_state', value=1, labels={'state':response['state']})
    yield metric

    for mtr in stats_metrics:
        mtype = 'gauge'
        if mtr in ['num_requests','num_failures']:
            mtype = 'counter'
        metric = Metric('locust_requests_'+mtr, 'Locust requests '+mtr, mtype)
        for stat in response['stats']:
            if not 'Total' in stat['name']:
                if mtr in stat:
                    metric.add_sample('locust_requests_'+mtr, value=stat[mtr], labels={'path':stat['name'], 'method':stat['method']})
                else:
                    print(mtr)
        yield metric

if __name__ == '__main__':
  # Usage: locust_exporter.py <port> <locust_host:port>
    try:
        start_http_server(int(os.environ.get('LISTENER_PORT', 8080)))
        REGISTRY.register(LocustCollector(os.environ['LOCUST']))
        print("Connecting to locust on: " + os.environ['LOCUST'])
        while True: time.sleep(1000)
    except KeyboardInterrupt:
        exit(0)
