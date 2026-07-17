package partials

import "os"

func backendURL() string {
	if u := os.Getenv("BACKEND_URL"); u != "" {
		return u
	}
	return "http://localhost:8080"
}
