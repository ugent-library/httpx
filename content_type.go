// Package httpx contains convenience HTTP helper functions.
package httpx

import (
	"io"
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

func DetectContentType(f io.ReadSeeker) (string, error) {
	b := make([]byte, 512)
	if _, err := f.Read(b); err != nil {
		return "", err
	}

	mimetype := http.DetectContentType(b)

	// rewind
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	return mimetype, nil
}
