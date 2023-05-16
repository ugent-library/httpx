package httpx

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

var (
	JSONContentType = "application/json"
	HTMLContentType = "text/html; charset=utf-8"
)

var bufPool = sync.Pool{
	New: func() any { return &bytes.Buffer{} },
}

func RenderJSON(w http.ResponseWriter, status int, v any) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufPool.Put(buf)
	}()

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if hdr := w.Header(); hdr.Get("Content-Type") == "" {
		hdr.Set("Content-Type", JSONContentType)
	}

	w.WriteHeader(status)

	w.Write(buf.Bytes())
}

func RenderHTML(w http.ResponseWriter, status int, v string) {
	if hdr := w.Header(); hdr.Get("Content-Type") == "" {
		hdr.Set("Content-Type", HTMLContentType)
	}

	w.WriteHeader(status)

	w.Write([]byte(v))
}
