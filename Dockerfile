FROM scratch
EXPOSE 9646
USER 1000

COPY locust_exporter /bin/locust_exporter

ENTRYPOINT ["locust_exporter"]
