package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":3000", nil)
}

// func healthHandler(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Add("Content-Type", "text/html")
// 	fmt.Fprintf(rw, "<html><body>Requested URL: %s</body></html>", r.URL.Path)
// }

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	bs, err := json.Marshal(HealthCheck{
		Status: "Ok",
	})
	if err != nil {
		fmt.Fprintf(rw, "marshal error")
		return
	}

	rw.Write(bs)
	// io.WriteString(rw, string(bs))
}
