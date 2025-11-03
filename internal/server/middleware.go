package server

import (
	"log"
	"net/http"
	"time"
)

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

/// Middleware para logar as requisições saindo
func (s *Server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		now := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s in %s", r.Method, r.URL.Path, time.Since(now))
	})
}
