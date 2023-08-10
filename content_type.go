// Package httpx contains convenience HTTP helper functions.
package httpx

import (
	"mime"
	"net/http"
	"strings"
)

func HasContentType(r *http.Request, mimetypes ...string) bool {
	ct := r.Header.Get("Content-Type")
	for _, v := range strings.Split(ct, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		for _, mt := range mimetypes {
			if t == mt {
				return true
			}
		}
	}
	return false
}
