# locust_exporter
A locust exporter for prometheus

This is a simple exporter for http://locust.io metrics. You get all the necessary details about current tests and the state of the locust.

Errors and requests stats are added with the method and path labels - BE CAREFUL - if you have a lot of endpoints. It is probably better to group the endpoints in your locustfile (please see: http://docs.locust.io/en/latest/writing-a-locustfile.html#grouping-requests-to-urls-with-dynamic-parameters).

Requirements: prometheus_client (sudo pip install prometheus_client)

Running the exporter:

`<LISTENER_PORT=1234> <LOCUST=localhost:8089> python ./locust_exporter.py`

i.e.:

`LISTENER_PORT=1234 LOCUST=localhost:8089 python ./locust_exporter.py`

![](https://github.com/mbolek/locust_exporter/blob/master/locust_exporter.png)
