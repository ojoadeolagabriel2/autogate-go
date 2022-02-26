package middlewares

import "net/http"

func CommonMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-App", "autogate-go")
		handler.ServeHTTP(w, r)
	})
}