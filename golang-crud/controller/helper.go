package controller

import (
	"encoding/json"
	"net/http"
	"strings"
)

func writeJSON(
	w http.ResponseWriter,
	status int,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func getIDFromPath(r *http.Request) string {
	return strings.TrimPrefix(
		r.URL.Path,
		"/users/",
	)
}
