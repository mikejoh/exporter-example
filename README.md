# Prometheus Exporter example

_Work in progress!_

This is an example of how to write a Prometheus Exporter in Go.

In this example we'll explore how to write an exporter using Go and a simple service that exposes a REST API. The service API can be seen as an third party application that exposes some kind of data (of it's internal state) that could be exposed to Prometheus. Instead of rewriting the application itself to export metrics in the Prometheus format we'll create a exporter that does this for us.
 
There are so many different ways (flavors) to do this when it comes to the details of an exporter:

* Should you use the standard library package for handling command line flags or something like `kingpin`?
* The naming of the CLI flags
* How should you log messages, using the `log` package provided by the Prometheus client package or something like `logrus`?
* Should you expose a link to the `/metrics` path when hitting `/`?

I want to keep this example as simple as possible, in the end it depends on how your third party application exposes data that can be exported to Prometheus. Is it a REST API? Is it a binary that you use to query a system or similar?

### Run this example

1. Run both Docker containers
```
docker build ./service-api -t service-api:latest
docker build ./service-api-exporter -t service-api-exporter:latest
docker network create service-api
docker run -d --name service-api --network service-api -p 8000:8000 service-api:latest
docker run -d --name service-api-exporter --network service-api -p 9100:9100 service-api-exporter:latest -service-api.url http://service-api:8000/api/info
```
2. Increment the items via the REST API:
```
curl -H "Content-Type: application/json" -X POST -d '{"number": 10}' http://localhost:8000/api/items
```
3. Browse to the exporter to see the exported metric: `http://localhost:9100/metrics`