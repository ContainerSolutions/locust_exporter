#!/usr/bin/python
# -*- coding: UTF-8 -*-
# Requires prometheus_client library:
# sudo pip install prometheus_client
from prometheus_client import start_http_server, Metric, REGISTRY
import json
import requests
import sys
import time

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
        exit(2)

    response = json.loads(response)

    stats_metrics=['avg_content_length','avg_response_time','current_rps','max_response_time','median_response_time','min_response_time','num_failures','num_requests']

    metric = Metric('locust_user_count', 'Swarmed users', 'gauge')
    metric.add_sample('locust_user_count', value=response['user_count'], labels={})
    yield metric

    metric = Metric('locust_errors', 'Locust requests errors', 'gauge')
    for err in response['errors']:
        metric.add_sample('locust_errors', value=err['occurences'], labels={'path':err['name'], 'method':err['method']})
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
                metric.add_sample('locust_requests_'+mtr, value=stat[mtr], labels={'path':stat['name'], 'method':stat['method']})
        yield metric

if __name__ == '__main__':
  # Usage: locust_exporter.py <port> <locust_host:port>
  if len(sys.argv) != 3:
      print('Usage: locust_exporter.py <port> <locust_host:port>')
      exit(1)
  else:
    try:
        start_http_server(int(sys.argv[1]))
        REGISTRY.register(LocustCollector(str(sys.argv[2])))
        print("Connecting to locust on: " + sys.argv[2])
        while True: time.sleep(1)
    except KeyboardInterrupt:
        exit(0)
