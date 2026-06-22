package observability

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func TracedTransport() http.RoundTripper {
	return otelhttp.NewTransport(http.DefaultTransport)
}

func TracedHTTPClient() *http.Client {
	return &http.Client{
		Transport: TracedTransport(),
	}
}
