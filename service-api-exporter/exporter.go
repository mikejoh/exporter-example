package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const (
	namespace = "service_api"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last query of Service API successful.",
		nil, nil,
	)
	numOfItems = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "items"),
		"How many items are stored in the Service API",
		nil, nil,
	)
)

type Info struct {
	NumItems int `json:"num_items"`
}

type Exporter struct {
	client *http.Client
	apiUri string
}

func NewExporter(uri string) (*Exporter, error) {
	h := &http.Client{Timeout: 10 * time.Second}

	return &Exporter{
		client: h,
		apiUri: uri,
	}, nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- numOfItems
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	numItems, err := e.fetchNumberOfItems()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		log.Errorf("Can't query Service API: %v", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		numOfItems, prometheus.GaugeValue, float64(numItems),
	)
}

func (e *Exporter) fetchNumberOfItems() (int, error) {
	r, err := e.client.Get(e.apiUri)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	var info Info

	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		return 0, err
	}

	return info.NumItems, nil
}
