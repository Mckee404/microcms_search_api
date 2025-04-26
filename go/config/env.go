package config

import (
	"os"
)

func GetApiKey() string {
	apiKey := os.Getenv("MICROCMS_API_KEY")
	if apiKey == "" {
		panic("MICROCMS_API_KEY is not set")
	}
	return apiKey
}

func GetServiceDomain() string {
	serviceDomain := os.Getenv("MICROCMS_SERVICE_DOMAIN")
	if serviceDomain == "" {
		panic("MICROCMS_SERVICE_DOMAIN is not set")
	}
	return serviceDomain
}