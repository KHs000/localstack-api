package v1

import "net/http"

// POST TODO
func POST(r *http.Request) bool {
	if r.Method != http.MethodPost {
		return false
	}
	return true
}

// GET TODO
func GET(r *http.Request) bool {
	if r.Method != http.MethodGet {
		return false
	}
	return true
}
