package observability

import "log"

// Init adalah single entry point untuk semua komponen observability.
// Cukup panggil observability.Init() di main.go — tidak perlu ubah kode lain.
func Init() {
	serviceName := GetServiceName()
	log.Printf("Initializing observability for service: %s\n", serviceName)

	if err := initTracing(); err != nil {
		log.Printf("Tracing init failed: %v", err)
	} else {
		log.Println("Tracing initialized → Alloy → Tempo")
	}

	if err := initProfiling(); err != nil {
		log.Printf("Profiling init failed: %v", err)
	} else {
		log.Println("Profiling initialized → Pyroscope")
	}

	initMetrics()
	log.Println("Observability initialized")
}
