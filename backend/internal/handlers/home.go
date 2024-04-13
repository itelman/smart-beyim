package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello man")
}
