package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendData(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	output, _ := json.Marshal(&data)
	fmt.Fprintln(w, string(output))
}

func sendError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	fmt.Fprintln(w, "Resource not found")
}
