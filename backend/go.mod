module github.com/lgtm-fp/ecommerce-backend

go 1.22

require (
	github.com/grafana/pyroscope-go v1.1.1
	github.com/lib/pq v1.10.9
	github.com/prometheus/client_golang v1.19.1
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.3.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
	go.opentelemetry.io/otel/trace v1.27.0
)
