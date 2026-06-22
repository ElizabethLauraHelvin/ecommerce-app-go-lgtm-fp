package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// HTTP Metrics
var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
		Namespace: "ecommerce",
	}, []string{"service", "method", "path", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request duration in seconds",
		Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5},
		Namespace: "ecommerce",
	}, []string{"service", "method", "path"})
)

// Database Metrics
var (
	dbQueryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "db_query_duration_seconds",
		Help:    "Database query duration in seconds",
		Buckets: []float64{.001, .005, .01, .05, .1, .5, 1, 5},
		Namespace: "ecommerce",
	}, []string{"service", "operation"})

	dbErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "db_errors_total",
		Help: "Total database errors",
		Namespace: "ecommerce",
	}, []string{"service", "operation"})
)

// Business Metrics - Product
var (
	productsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "products_total",
		Help: "Total products in system",
		Namespace: "ecommerce",
	}, []string{"service"})

	productQueriesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "product_queries_total",
		Help: "Total product queries",
		Namespace: "ecommerce",
	}, []string{"service", "status"})
)

// Business Metrics - Orders
var (
	ordersTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "orders_total",
		Help: "Total orders created",
		Namespace: "ecommerce",
	}, []string{"service"})

	orderValue = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "order_value_usd",
		Help:    "Order value in USD",
		Buckets: []float64{10, 50, 100, 250, 500, 1000, 5000},
		Namespace: "ecommerce",
	}, []string{"service"})

	orderItems = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "order_items_count",
		Help:    "Number of items in order",
		Buckets: []float64{1, 2, 5, 10, 25},
		Namespace: "ecommerce",
	}, []string{"service"})
)

// Service Metrics
var (
	serviceVersion = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "service_version",
		Help: "Service version info",
		Namespace: "ecommerce",
	}, []string{"service", "version"})
)

func initMetrics() {
	serviceName := GetServiceName()
	version := GetEnv("APP_VERSION", "1.0.0")
	serviceVersion.WithLabelValues(serviceName, version).Set(1)
}

// Recording functions
type MetricsRecorder struct {
	serviceName string
}

func NewMetricsRecorder(serviceName string) *MetricsRecorder {
	return &MetricsRecorder{serviceName: serviceName}
}

func (m *MetricsRecorder) RecordHTTPRequest(method, path, status string, duration float64) {
	httpRequestsTotal.WithLabelValues(m.serviceName, method, path, status).Inc()
	httpRequestDuration.WithLabelValues(m.serviceName, method, path).Observe(duration)
}

func (m *MetricsRecorder) RecordDBQuery(operation string, duration float64, err error) {
	dbQueryDuration.WithLabelValues(m.serviceName, operation).Observe(duration)
	if err != nil {
		dbErrors.WithLabelValues(m.serviceName, operation).Inc()
	}
}

func (m *MetricsRecorder) RecordProductQuery(status string) {
	productQueriesTotal.WithLabelValues(m.serviceName, status).Inc()
}

func (m *MetricsRecorder) RecordProductTotal(count float64) {
	productsTotal.WithLabelValues(m.serviceName).Set(count)
}

func (m *MetricsRecorder) RecordOrder(itemCount float64, totalValue float64) {
	ordersTotal.WithLabelValues(m.serviceName).Inc()
	orderItems.WithLabelValues(m.serviceName).Observe(itemCount)
	orderValue.WithLabelValues(m.serviceName).Observe(totalValue)
}
