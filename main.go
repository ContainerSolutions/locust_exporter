package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	namespace = "locust"
)

// Exporter structure
type Exporter struct {
	uri   string
	mutex sync.RWMutex
	fetch func(endpoint string) (io.ReadCloser, error)

	locustUp, locustUsers, locustSlaves, locustFailRatio, locustCurrentResponseTimePercentileNinetyFifth, locustCurrentResponseTimePercentileFiftieth                                                                                                                                                        prometheus.Gauge
	locustNumRequests, locustNumFailures, locustAvgResponseTime, locustNinetiethResponseTime, locustCurrentFailPerSec, locustMinResponseTime, locustMaxResponseTime, locustCurrentRps, locustMedianResponseTime, locustAvgContentLength, locustErrors *prometheus.GaugeVec
	totalScrapes                                                                                                                                                                                prometheus.Counter
}

// NewExporter function
func NewExporter(uri string, timeout time.Duration) (*Exporter, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	var fetch func(endpoint string) (io.ReadCloser, error)
	switch u.Scheme {
	case "http", "https", "file":
		fetch = fetchHTTP(uri, timeout)
	default:
		return nil, fmt.Errorf("unsupported scheme: %q", u.Scheme)
	}

	return &Exporter{
		uri:   uri,
		fetch: fetch,
		locustUp: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "up",
				Help:      "The current health status of the server (1 = UP, 0 = DOWN).",
			},
		),
		locustUsers: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "users",
				Help:      "The current number of users.",
			},
		),
		locustSlaves: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "slaves",
				Help:      "The current number of slaves.",
			},
		),
		locustCurrentResponseTimePercentileNinetyFifth: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "current_response_time_percentile_95",
			},
		),
		locustCurrentResponseTimePercentileFiftieth: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "current_response_time_percentile_50",
			},
		),
		locustFailRatio: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "fail_ratio",
			},
		),
		locustNumRequests: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "num_requests",
			},
			[]string{"method", "name"},
		),
		locustNumFailures: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "num_failures",
			},
			[]string{"method", "name"},
		),
		locustAvgResponseTime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "avg_response_time",
			},
			[]string{"method", "name"},
		),
		locustNinetiethResponseTime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "ninetieth_response_time",
			},
			[]string{"method", "name"},
		),
		locustCurrentFailPerSec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "current_fail_per_sec",
			},
			[]string{"method", "name"},
		),
		locustMinResponseTime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "min_response_time",
			},
			[]string{"method", "name"},
		),
		locustMaxResponseTime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "max_response_time",
			},
			[]string{"method", "name"},
		),
		locustCurrentRps: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "current_rps",
			},
			[]string{"method", "name"},
		),
		locustMedianResponseTime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "median_response_time",
			},
			[]string{"method", "name"},
		),
		locustAvgContentLength: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "requests",
				Name:      "avg_content_length",
			},
			[]string{"method", "name"},
		),
		locustErrors: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "errors",
				Help:      "The current number of errors.",
			},
			[]string{"method", "name"},
		),
		totalScrapes: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "total_scrapes",
				Help:      "The total number of scrapes.",
			},
		),
	}, nil
}

// Describe function of Exporter
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	ch <- e.locustUsers.Desc()
	ch <- e.locustSlaves.Desc()
	ch <- e.locustUp.Desc()
	ch <- e.totalScrapes.Desc()
	ch <- e.locustFailRatio.Desc()
	ch <- e.locustCurrentResponseTimePercentileNinetyFifth.Desc()
	ch <- e.locustCurrentResponseTimePercentileFiftieth.Desc()

	e.locustNumRequests.Describe(ch)
	e.locustNumFailures.Describe(ch)
	e.locustAvgResponseTime.Describe(ch)
	e.locustNinetiethResponseTime.Describe(ch)
	e.locustCurrentFailPerSec.Describe(ch)
	e.locustMinResponseTime.Describe(ch)
	e.locustMaxResponseTime.Describe(ch)
	e.locustMedianResponseTime.Describe(ch)
	e.locustCurrentRps.Describe(ch)
	e.locustAvgContentLength.Describe(ch)
	e.locustErrors.Describe(ch)
}

// Collect function of Exporter
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	up := e.scrape(ch)
	ch <- prometheus.MustNewConstMetric(e.locustUp.Desc(), prometheus.GaugeValue, up)
	e.locustNumRequests.Collect(ch)
	e.locustNumFailures.Collect(ch)
	e.locustAvgResponseTime.Collect(ch)
    e.locustNinetiethResponseTime.Collect(ch)
    e.locustCurrentFailPerSec.Collect(ch)
	e.locustMinResponseTime.Collect(ch)
	e.locustMaxResponseTime.Collect(ch)
	e.locustCurrentRps.Collect(ch)
	e.locustMedianResponseTime.Collect(ch)
	e.locustAvgContentLength.Collect(ch)
	e.locustErrors.Collect(ch)
}

type locustStats struct {
	Stats []struct {
		Method                  string  `json:"method"`
		Name                    string  `json:"name"`
		NumRequests             int     `json:"num_requests"`
		NumFailures             int     `json:"num_failures"`
		AvgResponseTime         float64 `json:"avg_response_time"`
		NinetiethResponseTime   float64 `json:"ninetieth_response_time"`
		CurrentFailPerSec       float64 `json:"current_fail_per_sec"`
		MinResponseTime         float64 `json:"min_response_time"`
		MaxResponseTime         float64 `json:"max_response_time"`
		CurrentRps              float64 `json:"current_rps"`
		MedianResponseTime      float64 `json:"median_response_time"`
		AvgContentLength        float64 `json:"avg_content_length"`
	} `json:"stats"`
	Errors []struct {
		Method     string `json:"method"`
		Name       string `json:"name"`
		Error      string `json:"error"`
		Occurences int    `json:"occurences"`
	} `json:"errors"`
	TotalRps                                    float64 `json:"total_rps"`
	FailRatio                                   float64 `json:"fail_ratio"`
	CurrentResponseTimePercentileNinetyFifth    float64 `json:"current_response_time_percentile_95"`
	CurrentResponseTimePercentileFiftieth       float64 `json:"current_response_time_percentile_50"`
	SlaveCount                                  int     `json:"slave_count,omitempty"`
	State                                       string  `json:"state"`
	UserCount                                   int     `json:"user_count"`
}

func (e *Exporter) scrape(ch chan<- prometheus.Metric) (up float64) {
	e.totalScrapes.Inc()

	var locustStats locustStats

	body, err := e.fetch("/stats/requests")
	if err != nil {
		log.Errorf("Can't scrape Pack: %v", err)
		return 0
	}
	defer body.Close()

	bodyAll, err := ioutil.ReadAll(body)
	if err != nil {
		return 0
	}

	_ = json.Unmarshal([]byte(bodyAll), &locustStats)

	ch <- prometheus.MustNewConstMetric(e.locustUsers.Desc(), prometheus.GaugeValue, float64(locustStats.UserCount))
	ch <- prometheus.MustNewConstMetric(e.locustSlaves.Desc(), prometheus.GaugeValue, float64(locustStats.SlaveCount))
	ch <- prometheus.MustNewConstMetric(e.locustFailRatio.Desc(), prometheus.GaugeValue, float64(locustStats.FailRatio))
	ch <- prometheus.MustNewConstMetric(e.locustCurrentResponseTimePercentileNinetyFifth.Desc(), prometheus.GaugeValue, float64(locustStats.CurrentResponseTimePercentileNinetyFifth))
	ch <- prometheus.MustNewConstMetric(e.locustCurrentResponseTimePercentileFiftieth.Desc(), prometheus.GaugeValue, float64(locustStats.CurrentResponseTimePercentileFiftieth))

	for _, r := range locustStats.Stats {
		if r.Name != "Total" && r.Name != "//stats/requests" {
			e.locustNumRequests.WithLabelValues(r.Method, r.Name).Set(float64(r.NumRequests))
			e.locustNumFailures.WithLabelValues(r.Method, r.Name).Set(float64(r.NumFailures))
			e.locustAvgResponseTime.WithLabelValues(r.Method, r.Name).Set(r.AvgResponseTime)
			e.locustNinetiethResponseTime.WithLabelValues(r.Method, r.Name).Set(r.NinetiethResponseTime)
			e.locustCurrentFailPerSec.WithLabelValues(r.Method, r.Name).Set(r.CurrentFailPerSec)
			e.locustMinResponseTime.WithLabelValues(r.Method, r.Name).Set(r.MinResponseTime)
			e.locustMaxResponseTime.WithLabelValues(r.Method, r.Name).Set(r.MaxResponseTime)
			e.locustCurrentRps.WithLabelValues(r.Method, r.Name).Set(r.CurrentRps)
			e.locustMedianResponseTime.WithLabelValues(r.Method, r.Name).Set(r.MedianResponseTime)
			e.locustAvgContentLength.WithLabelValues(r.Method, r.Name).Set(r.AvgContentLength)
		}
	}

	for _, r := range locustStats.Errors {
		e.locustErrors.WithLabelValues(r.Method, r.Name).Set(float64(r.Occurences))
	}

	return 1
}

func fetchHTTP(uri string, timeout time.Duration) func(endpoint string) (io.ReadCloser, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	return func(endpoint string) (io.ReadCloser, error) {
		resp, err := client.Get(uri + endpoint)
		if err != nil {
			return nil, err
		}
		if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
			resp.Body.Close()
			return nil, fmt.Errorf("HTTP status %d", resp.StatusCode)
		}
		return resp.Body, nil
	}
}

func main() {

	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9646").Envar("LOCUST_EXPORTER_WEB_LISTEN_ADDRESS").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").Envar("LOCUST_EXPORTER_WEB_TELEMETRY_PATH").String()
		uri           = kingpin.Flag("locust.uri", "URI of Locust.").Default("http://localhost:8089").Envar("LOCUST_EXPORTER_URI").String()
		timeout       = kingpin.Flag("locust.timeout", "Scrape timeout").Default("5s").Envar("LOCUST_EXPORTER_TIMEOUT").Duration()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("locust_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting locust_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	exporter, err := NewExporter(*uri, *timeout)
	if err != nil {
		log.Fatal(err)
	}
	prometheus.MustRegister(exporter)
	prometheus.MustRegister(version.NewCollector("locustexporter"))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html><head><title>Locust Exporter</title></head><body><h1>Locust Exporter</h1><p><a href='` + *metricsPath + `'>Metrics</a></p></body></html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}