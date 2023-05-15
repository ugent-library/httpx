package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { return &bytes.Buffer{} },
}

func RenderJSON(w io.Writer, v any) error {
	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufPool.Put(buf)
	}()

	enc := json.NewEncoder(buf)

	if err := enc.Encode(v); err != nil {
		return err
	}

	if hw, ok := w.(http.ResponseWriter); ok {
		if hw.Header().Get("Content-Type") == "" {
			hw.Header().Set("Content-Type", "application/json")
		}
	}

	w.Write(buf.Bytes())

	return nil
}
