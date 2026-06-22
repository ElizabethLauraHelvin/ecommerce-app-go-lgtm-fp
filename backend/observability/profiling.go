package observability

import "github.com/grafana/pyroscope-go"

// initProfiling menginisialisasi Grafana Pyroscope untuk continuous profiling.
func initProfiling() error {
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: GetEnv("OTEL_SERVICE_NAME", "ecommerce-backend"),
		ServerAddress:   GetEnv("PYROSCOPE_ENDPOINT", "http://pyroscope.monitoring.svc.cluster.local:4040"),
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
		},
	})
	return err
}

func initMetrics() {}
