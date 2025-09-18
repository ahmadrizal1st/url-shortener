package utils

import (
	"os"
	"strings"
)

func IsDifferentDomain(url string) bool {
	domain := os.Getenv("DOMAIN")

	if url == domain {
		return false
	}

	cleanUrl := strings.TrimPrefix(url, "http://")
	cleanUrl = strings.TrimPrefix(cleanUrl, "https://")
	cleanUrl = strings.TrimPrefix(cleanUrl, "www.")
	cleanUrl = strings.Split(cleanUrl, "/")[0]

	return cleanUrl != domain
}

func EnsureHttPPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}
