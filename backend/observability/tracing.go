package observability

import (
	"context"
	"database/sql"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// initTracing menginisialisasi OpenTelemetry tracing.
// Trace dikirim ke Alloy → Tempo via OTLP HTTP.
func initTracing() error {
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(
			GetEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://alloy.monitoring.svc.cluster.local:4318"),
		),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return err
	}

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(GetEnv("OTEL_SERVICE_NAME", "ecommerce-backend")),
			semconv.DeploymentEnvironment(GetEnv("APP_ENV", "production")),
		),
	)
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)

	// Set propagator untuk membaca W3C traceparent header dari FE (Faro)
	// Tanpa ini, trace FE dan BE tidak terhubung meski traceparent sudah dikirim
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, // baca/tulis W3C traceparent header
			propagation.Baggage{},      // baca/tulis W3C baggage header
		),
	)

	return nil
}

// OpenDBWithTracing membuka koneksi PostgreSQL dengan instrumentasi otelsql.
// Setiap query SQL otomatis menghasilkan span di Tempo (trace sampai level DB).
func OpenDBWithTracing(dsn string) (*sql.DB, error) {
	db, err := otelsql.Open(
		"postgres",
		dsn,
		otelsql.WithAttributes(
			semconv.DBSystemPostgreSQL,
		),
		otelsql.WithDBName(GetEnv("DB_NAME", "ecommerce")),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}