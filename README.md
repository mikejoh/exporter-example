# Prometheus Exporter example

This is an example of how to write a Prometheus Exporter in Go.

In this example we'll explore how to write an exporter using Go and a simple service that exposes a REST API. The service API can be seen as an third party application that exposes some kind of data (of it's internal state) that could be exposed to Prometheus. Instead of rewriting the application itself to export metrics in the Prometheus format we'll create a exporter that does this for us.
 
There are so many different ways (flavors) to do this when it comes to the details of an exporter:

* Should you use the standard library package for handling command line flags or something like `kingpin`?
* The naming of the CLI flags
* How should you log messages?
* Should you expose a link to the `/metrics` path when hitting `/`, how can that be done?

I want to keep this example as simple as possible, in the end it depends on how your third party application exposes data that can be exported to Prometheus. Is it a REST API? Is it a binary that you use to query a system or similar?