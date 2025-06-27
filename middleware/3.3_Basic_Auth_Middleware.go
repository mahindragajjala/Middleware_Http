//3.3_Basic_Auth_Middleware
/* 
Basic Authentication is a simple HTTP authentication scheme in which:
The client sends a username and password in the request header.
The server decodes and verifies them before allowing access.
*/
package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func BasicAuthMiddleware(username, password string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		if err != nil {
			http.Error(w, "Invalid authentication", http.StatusBadRequest)
			return
		}

		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || pair[0] != username || pair[1] != password {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r) // Proceed to next
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, authenticated user!"))
}

func main() {
	username := "admin"
	password := "secret"

	protected := BasicAuthMiddleware(username, password, http.HandlerFunc(Hello))

	http.Handle("/", protected)
	http.ListenAndServe(":8080", nil)
}


/* 
curl -u admin:secret http://localhost:8080
# Output: Welcome, authenticated user!
*/
