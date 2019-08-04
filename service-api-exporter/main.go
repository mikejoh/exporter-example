package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
)

var (
	showVersion   = flag.Bool("version", false, "Print version information")
	listenAddress = flag.String("web.listen-address", ":9100", "Address to listen on for web interface and telemetry")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path to expose metrics of the exporter")
	serviceAPIURL = flag.String("service-api.url", "http://localhost:8000/api/info", "Address where to fetch the Service API info")
)

func init() {
	prometheus.MustRegister(version.NewCollector("service_api_exporter"))
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("service-api-exporter"))
		os.Exit(0)
	}

	if *serviceAPIURL == "" {
		fmt.Fprintln(os.Stderr, "Please provide a address for Service API")
		os.Exit(1)
	}

	log.Infoln("Starting Service API exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	e, err := NewExporter(*serviceAPIURL)
	if err != nil {
		fmt.Println("Error initializing Service API exporter.")
		os.Exit(1)
	}

	prometheus.MustRegister(e)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
	})

	log.Infof("Listening on %s", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
