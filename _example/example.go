package main

import (
	"net/http"
	"strings"

	"github.com/Pranc1ngPegasus/middlechain"
)

func main() {
	mux := http.NewServeMux()

	handler := middlechain.Chain(mux,
		heaatbeat("/ping"),
	)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}

func heaatbeat(endpoint string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if (r.Method == http.MethodGet || r.Method == http.MethodHead) && strings.EqualFold(r.URL.Path, endpoint) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("."))

				return
			}
		})
	}
}
