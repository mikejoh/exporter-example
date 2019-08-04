# Service API Exporter

_Work in progress!_

This exporter will collect and export metrics exposed by the Service API.

### Running the Exporter
```
docker build . -t service-api-exporter:latest
docker run --rm -p 9100:9100 service-api-exporter:latest
```