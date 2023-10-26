package httpx

import "net/http"

var (
	MethodOverrideHeader = "X-HTTP-Method-Override"
	MethodOverrideField  = "_method"
)

// The MethodOverride middleware checks for a HTTP method override in the request and
// uses it instead of the original method.
//
// For security reasons only `POST` can be overridden with `PUT`, `PATCH` or `DELETE`.
//
// MethodOverride is similar to the gorilla HTTPMethodOverrideHandler but checks
// for a X-HTTP-Method-Override header before the _method form field to avoid parsing
// the request body unnecessarily.
func MethodOverride(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			om := r.Header.Get(MethodOverrideHeader)
			if om == "" {
				om = r.FormValue(MethodOverrideField)
			}
			if om == "PUT" || om == "PATCH" || om == "DELETE" {
				r.Method = om
			}
		}
		h.ServeHTTP(w, r)
	})
}
