package api

import (
	"net/http"
	"os"

	"github.com/igilgyrg/arbitrage/api/middleware"
)

func httpInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func (s *Server) mapRoutes() {
	s.mux.Handle("/status", httpInterceptor(s.endpoints.Status()))

	apiKey := os.Getenv("CLIENT_API_KEY")
	if apiKey != "" {
		accessFunc := middleware.ClientAccess(s.logger, apiKey)
		s.mux.Handle("/bundles", httpInterceptor(accessFunc(s.endpoints.Bundles())))
	}
}
