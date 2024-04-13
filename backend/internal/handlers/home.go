package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	h.log.Info(r.URL.RawPath)
	fmt.Fprint(w, "hello man")
}
