# Service API

_Work in progress!_

We'll be using the REST API exposed by this simple service to collect the number of items stored in the service.

### Running the Service API
```
docker build . -t service-api:latest
docker run --rm -p 8000:8000 service-api:latest
```
