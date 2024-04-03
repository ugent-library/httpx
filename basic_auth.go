package httpx

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

// BasicAuth provides safe basic auth handling.
// See https://www.alexedwards.net/blog/basic-authentication-in-go.
func BasicAuth(username, password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if u, p, ok := r.BasicAuth(); ok {
				usernameHash := sha256.Sum256([]byte(u))
				passwordHash := sha256.Sum256([]byte(p))
				expectedUsernameHash := sha256.Sum256([]byte(username))
				expectedPasswordHash := sha256.Sum256([]byte(password))
				usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
				passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

				if usernameMatch && passwordMatch {
					h.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		})
	}
}
