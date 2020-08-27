package middleware

import (
	"net/http"
	"os"
)

// BasicAuth ...
func BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	basicAuthUsername := os.Getenv("BASIC_AUTH_USERNAME")
	basicAuthPassword := os.Getenv("BASIC_AUTH_PASSWORD")

	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`not authenticated`))
		return false
	}

	isValid := (username == basicAuthUsername) && (password == basicAuthPassword)
	if !isValid {
		w.Write([]byte(`wrong username or password`))
		return false
	}

	return true
}
