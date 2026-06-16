package configs

import "os"

type Config struct {
	Port                 string
	StatisticsServiceURL string
	StatisticsTimeoutMS  string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	timeout := os.Getenv("STATISTICS_TIMEOUT_MS")
	if timeout == "" {
		timeout = "5000"
	}

	return Config{
		Port:                 port,
		StatisticsServiceURL: os.Getenv("STATISTICS_SERVICE_URL"),
		StatisticsTimeoutMS:  timeout,
	}
}
