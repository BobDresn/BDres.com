package handlers

import (
	"net/http"
)

func HtmxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<p>This content was loaded via HTMX!</p>"))
}
