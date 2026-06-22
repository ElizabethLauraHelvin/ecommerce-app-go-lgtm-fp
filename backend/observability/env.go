package observability

import "os"

var serviceName string

// SetServiceName sets the service name for this instance
func SetServiceName(name string) {
	serviceName = name
}

// GetServiceName returns the current service name
func GetServiceName() string {
	if serviceName != "" {
		return serviceName
	}
	return GetEnv("OTEL_SERVICE_NAME", "ecommerce-service")
}

// GetEnv mengambil nilai environment variable dengan fallback ke defaultValue.
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
