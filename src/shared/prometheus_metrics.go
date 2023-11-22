package shared

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var APIRequestsMetrics = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "mini_url_service_api",
		Help: "Mini URL Service API HTTP requests metrics",
	}, []string{"request_url", "request_with_error", "http_method"},
)

var RedirectRequestsMetrics = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "mini_url_service_redirect",
		Help: "Mini URL Service HTTP redirects metrics",
	}, []string{"request_url", "request_with_error", "http_method"},
)
